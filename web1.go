package main

import (
    "net/http"
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Inside hanler")
    fmt.Fprintf(w, "Hello World from my Go Program!")
}

func main() {
    http.HandleFunc("/", handler) // 生定向所有URL到 handler 函数
    http.ListenAndServe("localhost:9999", nil) // 监听来自9999端口的连接请求
}
