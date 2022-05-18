package main

import (
	"fmt"
	"net"
	"time"
)

func main()  {
	ls, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("服务器监听成功")
	defer ls.Close()
	for  {
		conn, err := ls.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("接收到一个连接")
		processConn(conn)
	}
}

func processConn(conn net.Conn)  {
	defer conn.Close()
	buff := make([]byte, 5)
	for  {
		time.Sleep(time.Second)
		conn.SetReadDeadline(time.Now().Add(time.Second))
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("conn read %d bytes, error:%s \n", n, err)

			if e, ok := err.(net.Error); ok && e.Timeout() {
				fmt.Println("超时")
				continue
			}

			return
		}
		fmt.Printf("读取到: %s", string(buff[:n]))
	}
}