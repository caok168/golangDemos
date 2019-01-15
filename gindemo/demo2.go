package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	//生成要访问的url
	url := "http://localhost:9333/dir/assign"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	body, err_body := ioutil.ReadAll(response.Body)
	if err_body != nil {
		fmt.Println("err_body:", err_body)
	}
	fmt.Println(string(body))

	BodyValue := string(body)
	fmt.Println(BodyValue)
	var assign WeedAssign
	err = json.Unmarshal(body, &assign)
	if err != nil {
		fmt.Println("err:=", err)
	}

	req := `{"file":"@/home/ck/图片/111.JPG"}`
	//req_new := bytes.NewBuffer([]byte(req))

	postUrl := "http://" + assign.Url + "/" + assign.Fid

	res, err := DoBytesPost(postUrl, []byte(req))
	fmt.Println("res:", string(res))

	//request, err := http.NewRequest("POST", postUrl, req_new)
	//request.Header.Set("Content-type", "application/json")
	//if err != nil {
	//	fmt.Println("err:=", err)
	//}
	//
	//response1, err := client.Do(request)
	//if err != nil {
	//	fmt.Println("err:=", err)
	//}
	//if response1.StatusCode == 200 {
	//	body, _ := ioutil.ReadAll(response1.Body)
	//	fmt.Println(string(body))
	//}

}

//body提交二进制数据
func DoBytesPost(url string, data []byte) ([]byte, error) {

	body := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	//request.Header.Set("Connection", "Keep-Alive")
	request.Header.Set("Content-type", "multipart/form-data")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}

func test() {

	client := &http.Client{}
	req := `{"name":"junneyang", "age": 88}`
	req_new := bytes.NewBuffer([]byte(req))
	request, _ := http.NewRequest("POST", "http://10.67.2.252:8080/test/", req_new)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

type WeedAssign struct {
	Fid       string `json:"fid"`
	Url       string `json:"url"`
	PublicUrl string `json:"publicUrl"`
	Count     int    `json:"count"`
}
