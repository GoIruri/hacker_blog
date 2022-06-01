package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	//1 页面 views 2 api 数据(json) 3 静态资源
	//  http句柄，根目录和响应函数
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	//URL请求路径映射
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
