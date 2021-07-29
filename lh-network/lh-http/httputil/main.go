package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		return
	}
	d, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return
	}
	log.Println(string(d))

}
