package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 处理连接
	fmt.Println("Accepted connection from", conn.RemoteAddr())

	// 读取客户端消息
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// 处理接收到的消息
		message := string(buffer[:n])
		fmt.Println("Received message:", message)

		// 在这里可以编写处理消息的逻辑

		// 如果需要，可以向客户端发送响应
		// conn.Write([]byte("Message received"))
	}

	// 在这里可以编写与客户端通信的逻辑
}

func main() {

	// 启动Socket服务器
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Socket server listening on :8081")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 启动新的goroutine来处理连接
		go handleConnection(conn)
	}
}
