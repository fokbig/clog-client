package log

import (
	"LogCloud/pkg/setting"
	"log"
	"strings"
)

var FilenameArray []string

func Setup() {
	logSetting := setting.LogSetting
	path := logSetting.Path
	logList := logSetting.LogList
	FilenameArray = strings.Split(logList, ",")
	for i := range FilenameArray {
		FilenameArray[i] = path + "/" + FilenameArray[i]
	}
	log.Println("待上传日志列表: ")
	for _, logName := range FilenameArray {
		log.Println("日志名: " + logName)
	}
}
