package main

import (
	"fmt"
	"net"
)

func handleStatus(conn net.Conn) string {
	defer conn.Close()

	// 创建缓冲区
	buffer := make([]byte, 1024)

	// 读取接收到的文字数据
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取数据错误:", err)
	}
	// 提取文字内容
	status := string(buffer[:n])
	fmt.Println("接收到的文字:", status)
	return status
}

func dns() {
	listener, err := net.Listen("tcp", ":9888")
	if err != nil {
		fmt.Println("无法监听端口:", err)
	}
	defer listener.Close()

	fmt.Println("服务器已启动，等待连接...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("接受连接错误:", err)
	}
	//接受状态码，如果是同步区块链请求，则发送给该ip
	info := handleStatus(conn)
	status := info[0:3]
	ip := info[3:]

	if status == "bal" {
		//返回余额
		bal := getBalance(status[3:])
		_, err = conn.Write([]byte(string(bal)))
		if err != nil {
			fmt.Println("Error sending response:", err.Error())
			return
		}
		fmt.Printf("Responded %x balance request succed", status[3:])
	}
	if status == "syn" {
		//同步区块链
		sendBc(ip)
	}
	if status == "upd" {
		//收到来自节点的更新请求
		recvBc()
	}
}
