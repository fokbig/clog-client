package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Box struct {
	From     string   `json:"from"`
	Filename string   `json:"filename"`
	Lines    []string `json:"lines"`
}

// Post 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func (data *Box) Post(url string, contentType string) string {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Post 请求发送失败，跳过该盒子")
		return err.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Post 请求读取流关闭异常")
		}
	}(resp.Body)

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
