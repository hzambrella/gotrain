package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
)

func TestPrt(t *testing.T) {
	resp, err := http.Get("127.0.0.1:8080/clouldprt/prtorder?recId=250250")
	if err != nil {
		t.Fatal(err)
	}
	bresp, err := httputil.DumpResponse(resp, false)
	if err != nil {
		t.Fatal(resp)
	}
	fmt.Println(string(bresp))

}
