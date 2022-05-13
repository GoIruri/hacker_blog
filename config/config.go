package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

//定义一个结构来映射config.toml文件
type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	Username    string
	UserDesc    string
}

type SystemConfig struct {
	AppName        string
	Version        float32
	CurrentDir     string
	CdnURL         string
	QiniuAccessKey string
	QiniuSecretKey string
	Valine         bool
	ValineAppid    string
	ValineAppKey   string
	ValineAppURL   string
}

var Cfg *TomlConfig

func init() {
	//	程序运行时，就会执行init方法
	Cfg = &TomlConfig{}
	Cfg.System.AppName = "go-hacker-blog"
	Cfg.System.Version = 1.0
	Cfg.System.CurrentDir, _ = os.Getwd()
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err) //直接抛异常死掉
	}
}
