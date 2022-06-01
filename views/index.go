package views

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/models"
	"net/http"
)

//解析html
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index //拿到模板
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
	index.WriteData(w, hr) //填充数据

}
