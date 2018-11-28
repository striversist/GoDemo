package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}
	fmt.Fprint(resp, "Hello simple web!")
}

// 浏览器测试
// http://localhost:9090
// http://localhost:9090/?url_long=111&url_long=222
func main()  {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}