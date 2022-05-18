package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("连接服务器成功")
	for {
		time.Sleep(200 * time.Millisecond)
		buff := []byte("hello")
		n, err := conn.Write(buff)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("发送了%d个字节， 发送内容:%s", n, string(buff))
	}
}