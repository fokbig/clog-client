package util

import (
	"fmt"
	"net"
)

func Connect() net.Conn {
	conn, err := net.Dial("tcp", "localhost:10002")
	if err != nil {
		fmt.Println("TCP连接错误")
		return nil
	}
	return conn
}

func Send(conn net.Conn, data string) {
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("TCP发送失败")
	}
}
