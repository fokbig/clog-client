package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type App struct {
	From string
}

var AppSetting = &App{}

type Log struct {
	Path    string
	LogList string
}

var LogSetting = &Log{}

type Server struct {
	Address string
	Port    string
	LogPath string
}

var ServerSetting = &Server{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, 解析配置文件错误 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("log", LogSetting)
	mapTo("server", ServerSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
