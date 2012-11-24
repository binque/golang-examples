package main

import (
    "net/http"
    "fmt"
    "strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    remPartOfURL := r.URL.Path[len("/hello/"):] //获取 URL 地址中 /hello/ 部分之后的所有内容
    fmt.Fprintf(w, "Hello %s", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
    remPartOfURL := r.URL.Path[len("/shouthello/"):] // 同上
    fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

func main() {
    http.HandleFunc("/hello/", helloHandler)
    http.HandleFunc("/shouthello/", shouthelloHandler)
    http.ListenAndServe("localhost:9999", nil)
}
