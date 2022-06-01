package common

import (
	"go-blog/config"
	"go-blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
}
