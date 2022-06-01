// Package  provides ...
package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

//定义一个结构体
type indexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	//	程序入口，一个项目只能有一个入口
	//	web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
