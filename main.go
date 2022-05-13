package main

import (
	"go-blog/config"
	"go-blog/models"
	"log"
	"net/http"
	"text/template"
)

//定义一个结构体
type indexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

//解析html
func indexHtml(w http.ResponseWriter, r *http.Request) {
	t := template.New("index")
	//	拿到html文件的路径
	path := config.Cfg.System.CurrentDir
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板进行解析
	index := path + "/template/index.html"
	home := path + "/template/home.html"
	head := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	pagination := path + "/template/layout/pagination.html"
	post := path + "/template/layout/post.html"
	t, _ = template.ParseFiles(index, home, head, footer, personal, pagination, post)
	//页面上涉及到的所以数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go-hacker-blog",
			Content:      "内容",
			UserName:     "黑客",
			ViewCount:    123,
			CreateAt:     "2022-05-13",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, hr)
}

func main() {
	//	程序入口，一个项目只能有一个入口
	//	web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//  http句柄，根目录和响应函数
	http.HandleFunc("/", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
