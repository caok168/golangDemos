package rtsp

import (
	"bufio"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	// Client to server for presentation and stream objects; recommended
	DESCRIBE = "DESCRIBE"
	// Bidirectional for client and stream objects; optional
	ANNOUNCE = "ANNOUNCE"
	// Bidirectional for client and stream objects; optional
	GET_PARAMETER = "GET_PARAMETER"
	// Bidirectional for client and stream objects; required for Client to server, optional for server to client
	OPTIONS = "OPTIONS"
	// Client to server for presentation and stream objects; recommended
	PAUSE = "PAUSE"
	// Client to server for presentation and stream objects; required
	PLAY = "PLAY"
	// Client to server for presentation and stream objects; optional
	RECORD = "RECORD"
	// Server to client for presentation and stream objects; optional
	REDIRECT = "REDIRECT"
	// Client to server for stream objects; required
	SETUP = "SETUP"
	// Bidirectional for presentation and stream objects; optional
	SET_PARAMETER = "SET_PARAMETER"
	// Client to server for presentation and stream objects; required
	TEARDOWN = "TEARDOWN"
)

type Request struct {
	Method        string
	URL           *url.URL
	Proto         string
	ProtoMajor    int
	ProtoMinor    int
	Header        http.Header
	ContentLength int
	Body          io.ReadCloser
}

func (r Request) String() string {
	s := fmt.Sprintf("%s %s %s/%d.%d\r\n", r.Method, r.URL, r.Proto, r.ProtoMajor, r.ProtoMinor)
	for k, v := range r.Header {
		for _, v := range v {
			s += fmt.Sprintf("%s: %s\r\n", k, v)
		}
	}
	s += "\r\n"
	if r.Body != nil {
		str, _ := ioutil.ReadAll(r.Body)
		s += string(str)
	}
	return s
}

//生成32位md5字串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// RTSP Basic 验证
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// RTSP Digest 验证
func digestAuth(digestHeader DigestHeader) string {

	response := getResponse(digestHeader, "ANSI")

	auth := ""
	auth += "username=\"" + digestHeader.Username + "\", " + "realm=\"" + digestHeader.Realm + "\", "
	auth += "uri=\"" + digestHeader.Uri + "\","
	auth += "nonce=\"" + digestHeader.Nonce + "\"," + " response=\"" + response + "\""

	return auth
}

// MD5: response = md5( password:nonce:md5(public_method:url)  );
// ANSI:response= md5( md5(username:realm:password):nonce:md5(public_method:url) );
// 根据密码的不同类型，计算不同的response
func getResponse(digestHeader DigestHeader, passType string) string {
	response := ""

	if passType == "ANSI" {
		a1 := digestHeader.Username + ":" + digestHeader.Realm + ":" + digestHeader.Password
		ha1 := Md5(a1)
		a2 := digestHeader.Method + ":" + digestHeader.Uri
		ha2 := Md5(a2)
		a3 := ha1 + ":" + digestHeader.Nonce + ":" + ha2
		response = Md5(a3)
	} else if passType == "MD5" {
		a1 := digestHeader.Method + ":" + digestHeader.Uri
		ha1 := Md5(a1)

		a2 := digestHeader.Password + ":" + digestHeader.Nonce + ":" + ha1
		response = Md5(a2)
	}

	return response
}

func NewRequest(method, urlStr, cSeq string, authenticateValue string, body io.ReadCloser) (*Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req := &Request{
		Method:     method,
		URL:        u,
		Proto:      "RTSP",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     map[string][]string{"CSeq": []string{cSeq}},
		Body:       body,
	}
	var user string
	var pass string

	if u.User != nil {
		user = u.User.Username()
		pass, _ = u.User.Password()
	}

	if authenticateValue == "" {
		req.Header.Add("Authorization", "Basic "+basicAuth(user, pass))
	} else {
		resArr := strings.Split(authenticateValue, ",")
		digestHeader := DigestHeader{}
		for _, re := range resArr {
			kv := strings.Split(re, "=")
			if strings.Trim(kv[0], " ") == "realm" {
				digestHeader.Realm = strings.Replace(kv[1], "\"", "", -1)
			}
			if strings.Trim(kv[0], " ") == "nonce" {
				digestHeader.Nonce = strings.Replace(kv[1], "\"", "", -1)
			}
			if strings.Trim(kv[0], " ") == "algorithm" {
				digestHeader.Algorithm = strings.Replace(kv[1], "\"", "", -1)
			}
		}

		digestHeader.Username = user
		digestHeader.Password = pass

		digestHeader.Uri = u.Scheme + "://" + u.Host // "rtsp://192.168.17.4:554"
		digestHeader.Method = method

		req.Header.Add("Authorization", "Digest "+digestAuth(digestHeader))
	}

	return req, nil
}

type Session struct {
	cSeq    int
	conn    net.Conn
	session string
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) nextCSeq() string {
	s.cSeq++
	return strconv.Itoa(s.cSeq)
}

func (s *Session) Describe(urlStr, authenticateValue string) (*Response, error) {
	req := &Request{}
	var err error
	if len(authenticateValue) > 7 {
		authType := authenticateValue[0:6]
		if authType == "Digest" {
			authenticateParams := authenticateValue[7:]
			req, err = NewRequest(DESCRIBE, urlStr, s.nextCSeq(), authenticateParams, nil)
		}
	} else {
		req, err = NewRequest(DESCRIBE, urlStr, s.nextCSeq(), "", nil)
	}

	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/sdp")

	if s.conn == nil {
		s.conn, err = net.Dial("tcp", req.URL.Host)
		if err != nil {
			return nil, err
		}
	}

	_, err = io.WriteString(s.conn, req.String())
	if err != nil {
		return nil, err
	}
	return ReadResponse(s.conn)
}

func (s *Session) Describe_(urlStr string) (*Response, error) {
	req, err := NewRequest(DESCRIBE, urlStr, s.nextCSeq(), "", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/sdp")

	if s.conn == nil {
		s.conn, err = net.Dial("tcp", req.URL.Host)
		if err != nil {
			return nil, err
		}
	}

	_, err = io.WriteString(s.conn, req.String())
	if err != nil {
		return nil, err
	}
	return ReadResponse(s.conn)
}

func (s *Session) Options(urlStr string) (*Response, error) {
	req, err := NewRequest(OPTIONS, urlStr, s.nextCSeq(), "", nil)
	if err != nil {
		panic(err)
	}

	if s.conn == nil {
		s.conn, err = net.Dial("tcp", req.URL.Host)
		if err != nil {
			return nil, err
		}
	}

	_, err = io.WriteString(s.conn, req.String())
	if err != nil {
		return nil, err
	}
	return ReadResponse(s.conn)
}

func (s *Session) Setup(urlStr, transport string) (*Response, error) {
	req, err := NewRequest(SETUP, urlStr, s.nextCSeq(), "", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Transport", transport)

	if s.conn == nil {
		s.conn, err = net.Dial("tcp", req.URL.Host)
		if err != nil {
			return nil, err
		}
	}

	_, err = io.WriteString(s.conn, req.String())
	if err != nil {
		return nil, err
	}
	resp, err := ReadResponse(s.conn)
	s.session = resp.Header.Get("Session")
	return resp, err
}

func (s *Session) Play(urlStr, sessionId string) (*Response, error) {
	req, err := NewRequest(PLAY, urlStr, s.nextCSeq(), "", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Session", sessionId)

	if s.conn == nil {
		s.conn, err = net.Dial("tcp", req.URL.Host)
		if err != nil {
			return nil, err
		}
	}

	_, err = io.WriteString(s.conn, req.String())
	if err != nil {
		return nil, err
	}
	return ReadResponse(s.conn)
}

func (s *Session) Teardown(urlStr, sessionId string) (*Response, error) {
	req, err := NewRequest(TEARDOWN, urlStr, s.nextCSeq(), "", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Session", sessionId)

	if s.conn == nil {
		s.conn, err = net.Dial("tcp", req.URL.Host)
		if err != nil {
			return nil, err
		}
	}

	_, err = io.WriteString(s.conn, req.String())
	if err != nil {
		return nil, err
	}
	return ReadResponse(s.conn)
}

type closer struct {
	*bufio.Reader
	r io.Reader
}

func (c closer) Close() error {
	if c.Reader == nil {
		return nil
	}
	defer func() {
		c.Reader = nil
		c.r = nil
	}()
	if r, ok := c.r.(io.ReadCloser); ok {
		return r.Close()
	}
	return nil
}

func ParseRTSPVersion(s string) (proto string, major int, minor int, err error) {
	parts := strings.SplitN(s, "/", 2)
	proto = parts[0]
	parts = strings.SplitN(parts[1], ".", 2)
	if major, err = strconv.Atoi(parts[0]); err != nil {
		return
	}
	if minor, err = strconv.Atoi(parts[0]); err != nil {
		return
	}
	return
}

// super simple RTSP parser; would be nice if net/http would allow more general parsing
func ReadRequest(r io.Reader) (req *Request, err error) {
	req = new(Request)
	req.Header = make(map[string][]string)

	b := bufio.NewReader(r)
	var s string

	// TODO: allow CR, LF, or CRLF
	if s, err = b.ReadString('\n'); err != nil {
		return
	}

	parts := strings.SplitN(s, " ", 3)
	req.Method = parts[0]
	if req.URL, err = url.Parse(parts[1]); err != nil {
		return
	}

	req.Proto, req.ProtoMajor, req.ProtoMinor, err = ParseRTSPVersion(parts[2])
	if err != nil {
		return
	}

	// read headers
	for {
		if s, err = b.ReadString('\n'); err != nil {
			return
		} else if s = strings.TrimRight(s, "\r\n"); s == "" {
			break
		}

		parts := strings.SplitN(s, ":", 2)
		req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	req.ContentLength, _ = strconv.Atoi(req.Header.Get("Content-Length"))
	req.Body = closer{b, r}
	return
}

type Response struct {
	Proto      string
	ProtoMajor int
	ProtoMinor int

	StatusCode int
	Status     string

	ContentLength int64

	Header http.Header
	Body   io.ReadCloser
}

func (res Response) String() string {
	s := fmt.Sprintf("%s/%d.%d %d %s\n", res.Proto, res.ProtoMajor, res.ProtoMinor, res.StatusCode, res.Status)
	for k, v := range res.Header {
		for _, v := range v {
			s += fmt.Sprintf("%s: %s\n", k, v)
		}
	}
	return s
}

func ReadResponse(r io.Reader) (res *Response, err error) {
	res = new(Response)
	res.Header = make(map[string][]string)

	b := bufio.NewReader(r)
	var s string

	// TODO: allow CR, LF, or CRLF
	if s, err = b.ReadString('\n'); err != nil {
		return
	}

	parts := strings.SplitN(s, " ", 3)
	res.Proto, res.ProtoMajor, res.ProtoMinor, err = ParseRTSPVersion(parts[0])
	if err != nil {
		return
	}

	if res.StatusCode, err = strconv.Atoi(parts[1]); err != nil {
		return
	}

	res.Status = strings.TrimSpace(parts[2])

	// read headers
	for {
		if s, err = b.ReadString('\n'); err != nil {
			return
		} else if s = strings.TrimRight(s, "\r\n"); s == "" {
			break
		}

		parts := strings.SplitN(s, ":", 2)
		res.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	res.ContentLength, _ = strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)

	res.Body = closer{b, r}
	return
}

//realm="Embedded Net DVR", nonce="152883479", algorithm="MD5"
type DigestHeader struct {
	Realm     string `json:"realm"`
	Nonce     string `json:"nonce"`
	Algorithm string `json:"algorithm"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Uri       string `json:"uri"`
	Response  string `json:"response"`
	Method    string `json:"method"`
}
