package main

import (
	"fmt"
	"net/http"
	"text/template"
	"code.google.com/p/google-api-go-client/urlshortener/v1" // 导入GOOGLE API包
)

// 本项目的网页模板
var tmpl = template.Must(template.New("URL Shortener Template").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8" /><title>Goo.gl网址缩短服务</title>
<style>
input, button { border: .1em solid #ccc; border-radius: .3em; padding: .5em; }
button { background-color: #f0f0f0; }
button:hover { background-color: #ccc;}
</style>
</head>
<body>
<h1>Goo.gl网址缩短服务</h1>
{{if .}}{{.}}<br /><hr /><br />{{end}}
<form action="/shorten" type="POST">
<label for="long-url">长网址：</label><input type="text" name="url" id="long-url" placeholder="请在这里输入您要缩短的网址" /><button><span>给我短址</span></button>
</form>
<br /><hr /><br />
<form action="/lengthen" type="POST">
<label for="short-url">短网址：</label><input type="text" name="url" id="short-url" placeholder="请在这里输入您要获取原始网址的短网址" /><button><span>给我长址</span></button>
</form></body></html>
`))

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url") // 获取由网页提交的网址
	svc, _ := urlshortener.New(http.DefaultClient) // 使用http中的default client创建一个新的 urlshortener 实例
	shorturl, _ := svc.Url.Insert(&urlshortener.Url { LongUrl: url, }).Do() // 填充长的网址然后呼叫缩短服务
	tmpl.Execute(w, fmt.Sprintf("<h2 class=\"url\"><a href=\"%s\">%s</a></h2><h3 class=\"url\">源始长网址为：<em>%s</em></h3>", shorturl.Id, shorturl.Id, url))
}

func handleLengthen(w http.ResponseWriter, r *http.Request) {
	url := "http://goo.go/" + r.FormValue("url")
	svc, _ := urlshortener.New(http.DefaultClient)
	longurl, err := svc.Url.Get(url).Do()
	if err != nil {
		fmt.Println("error: %v", err)
		return
	}
	tmpl.Execute(w, fmt.Sprintf("<h2 class=\"url\"><a href=\"%s\">%s</a></h2><h3 class=\"url\">短网址为：<em>%s</em></h3>", url, url, longurl))
}

func main() {
        http.HandleFunc("/", handleRoot)
        http.HandleFunc("/shorten", handleShorten)
        http.HandleFunc("/lengthen", handleLengthen)

        http.ListenAndServe(":8001", nil)
}

