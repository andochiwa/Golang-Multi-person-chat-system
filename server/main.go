package main

import (
	"fmt"
	"net"
)

func init() {
	address := "localhost:6379"
	initPool(address)
	initUserDao()
	initMessageDao()
}

func main() {

	fmt.Println("服务器在9876端口开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:9876")
	if err != nil {
		fmt.Println("net listen err =", err)
		return
	}

	// 监听成功就等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err =", err)
		}

		// 连接成功后启动协程和客户端保持通讯
		go process(conn)

	}

}
