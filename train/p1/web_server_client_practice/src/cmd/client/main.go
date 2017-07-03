package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	// server.go
	setHost   string     = "http://localhost"
	setPort   string     = ":8080"
	setPath   string     = "/test/modify"
	urlValues url.Values = url.Values{"key": {"test"}, "data": {"哈哈。。。"}}

)
//TODO: pt
func main() {
	u:=urlValues
	murl :=setHost +setPort +setPath
	tm1 := time.Now()

	//fmt.Println(fmt.Sprintf("post:%surl\n data:%s\n",murl,u))
	resp, err := http.PostForm(murl, u)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	tm2:=time.Now()
	fmt.Println(string(result),tm2.Sub(tm1))
}

