package main

import (
	"LogCloud/pkg/log"
	"LogCloud/pkg/setting"
	"LogCloud/service/read_service"
	"sync"
)

func init() {
	setting.Setup()
	log.Setup()
}

func main() {
	var wg sync.WaitGroup
	var count = len(log.FilenameArray)
	wg.Add(count)

	for _, filename := range log.FilenameArray {
		logName := filename
		go func() {
			read_service.ListenerLog(logName)
			wg.Done()
		}()
	}

	wg.Wait()
}
