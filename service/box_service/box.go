package box_service

import (
	"LogCloud/pkg/setting"
	"LogCloud/pkg/util"
	"container/list"
	"fmt"
	"time"
)

// Box 日志消息盒
type Box struct {
	From     string     `json:"from"`
	Filename string     `json:"filename"`
	Lines    *list.List `json:"lines"`
}

var boxMap = make(map[string]*Box)

// Push 加入日志信息
// 1. 往盒子中添加数据行
// 2. 检查盒子是否符合发送条件
// 	  2.1 符合:    发送盒子到日志中心, 移除盒子
//	  2.2 不符合:  无操作
func Push(filename string, message string) {
	pushBox(filename, message)
	isSend := checkBox(filename)
	if isSend {
		sendBox(filename)
		removeBox(filename)
	}
}

// pushBox 往盒子中添加数据行
func pushBox(filename string, message string) {
	box := boxMap[filename]
	if box == nil {
		box = &Box{
			From:     setting.AppSetting.From,
			Filename: filename,
			Lines:    list.New(),
		}
		boxMap[filename] = box
	}
	box.Lines.PushBack(message)
}

// checkBox 检查盒子是否符合发送条件
func checkBox(filename string) bool {
	box := boxMap[filename]
	lines := box.Lines
	return lines.Len() >= 6
}

// sendBox 发送盒子到日志中心
func sendBox(filename string) {
	startTime := time.Now().UnixNano() / 1e3
	box := boxMap[filename]
	serverSetting := setting.ServerSetting
	lines := box.Lines
	lineArray := make([]string, lines.Len())
	var data string
	i := 0
	for line := lines.Front(); line != nil; line = line.Next() {
		lineArray[i] = line.Value.(string)
		data = data + lineArray[i]
		i++
	}

	request := util.Box{
		From:     box.From,
		Filename: box.Filename,
		Lines:    lineArray,
	}
	url := serverSetting.Address + ":" + serverSetting.Port + serverSetting.LogPath
	_ = request.Post(url, "application/json")
	endTime := time.Now().UnixNano() / 1e3
	fmt.Println("盒子提交成功：", endTime-startTime, "us")
}

// removeBox 移除发送过的盒子
func removeBox(filename string) {
	boxMap[filename] = nil
}
