package main

import (
	"fmt"
	"gowebapi/conf"
	"gowebapi/service"
	"net/http"
)

func main() {
	//http://127.0.0.1:8000/go
	// 单独写回调函数
	conf.Init()
	http.HandleFunc("/go", myHandler)
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	// switch r.Header.Get("Content-Type") {}

	service.InsertMsg("test", "trainers", "bob", 2, "city")
	fmt.Println(service.FindMsg("test", "trainers", "Ash", 1, "city"))
	w.Write([]byte("www.5lmh.com"))

}
