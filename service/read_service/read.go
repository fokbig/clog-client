package read_service

// tail命令用途是依照要求将指定的文件的最后部分输出到标准设备，通常是终端，
// 通俗讲来，就是把某个档案文件的最后几行显示到终端上，
// 假设该档案有更新，tail会自己主动刷新，确保你看到最新的档案内容
// 在日志收集中可以实时的监测日志的变化，Log需要换行才可以表现
import (
	"LogCloud/service/box_service"
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func ListenerLog(filename string) {
	location := tail.SeekInfo{Offset: 0, Whence: 2}
	config := tail.Config{
		ReOpen:    true,      // 重新打开
		Follow:    true,      // 是否跟随
		Location:  &location, // 从文件的哪个地方开始读
		MustExist: false,     // 文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 开始读取数据
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines //遍历chan，读取日志内容
		if !ok {
			fmt.Printf("tail file close reopen,fileName:%s\n",
				tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		box_service.Push(filename, msg.Text)
	}
}
