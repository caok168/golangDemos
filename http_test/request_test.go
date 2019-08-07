package http_test

import (
	"fmt"
	. "net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {
	req := &Request{Method: "GET"}

	req.URL, _ = url.Parse("http://www.google.com/search?q=foo&q=bar")

	req.ParseForm()
	a := req.Form["q"]
	fmt.Println(a)

	for k, v := range req.Form {
		fmt.Printf("%s, %s", k, v[0])
	}

	if q := req.FormValue("q"); q != "foo" {
		t.Errorf(`req.FormValue("q") = %q, want "foo"`, q)
	}

	if qs := req.Form["q"]; !reflect.DeepEqual(qs, []string{"foo", "bar"}) {
		t.Errorf(`req.Form["q"] = %q, want ["foo","bar"]`, qs)
	} else {
		fmt.Println(qs)
	}
}

func TestParseFormQuery(t *testing.T) {
	req, _ := NewRequest("POST", "http://www.google.com/search?q=foo&q=bar&both=x&prio=1&orphan=nope&empty=not",
		strings.NewReader("z=post&both=y&prio=2&=nokey&orphan;empty=&"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if q := req.FormValue("q"); q != "foo" {
		t.Errorf(`req.FormValue("q") = %q, want "foo"`, q)
	}

	if z := req.FormValue("z"); z != "post" {
		t.Errorf(`req.FormValue("z") = %q, want "post"`, z)
	}

	if bq, found := req.PostForm["q"]; found {
		t.Errorf(`req.PostForm["q"] = %q,want no entry in map`, bq)
	} else {
		//fmt.Println(bq)
	}

	if bz := req.PostFormValue("z"); bz != "post" {
		t.Errorf(`req.PostFormValue("z") = %q, want "post"`, bz)
	} else {
		//fmt.Println(bz)
	}

	if qs := req.Form["q"]; !reflect.DeepEqual(qs, []string{"foo", "bar"}) {
		t.Errorf(`req.Form["q"] = %q, want ["foo","bar"]`, qs)
	} else {
		fmt.Println(qs)
	}

	req.ParseForm()

	both := req.Form["both"]
	fmt.Println(both)

	if prio := req.FormValue("prio"); prio != "2" {
		t.Errorf(`req.FormValue("prio") = %q, want "2" (from body)`, prio)
	}

	if empty := req.Form["empty"]; !reflect.DeepEqual(empty, []string{"", "not"}) {
		t.Errorf(`req.FormValue("empty") = %q, want "" (from body)`, empty)
	}

	if nokey := req.Form[""]; !reflect.DeepEqual(nokey, []string{"nokey"}) {
		t.Errorf(`req.FormValue("nokey") = %q, want "nokey" (from body)`, nokey)
	}

	//fmt.Println(req.Body)
	//
	//p, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//
	//fmt.Println("p:", string(p))

}
