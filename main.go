package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

//定义一个结构体
type indexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

//响应函数
func index(w http.ResponseWriter, r *http.Request) {
	var indexData indexData
	indexData.Title = "1mjtrack"
	indexData.Desc = "hacker"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

//解析html
func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData indexData
	indexData.Title = "1mjtrack"
	indexData.Desc = "hacker"
	t := template.New("index")
	//	拿到html文件的路径
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("path is error:", err)
		return
	}
	t, _ = template.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	//	程序入口，一个项目只能有一个入口
	//	web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//http句柄，根目录和响应函数
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
