package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello World")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("服務器啟動失敗")
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			continue
		}

		go handleConnection(accept)
	}
}

func handleConnection(accept net.Conn) {
	fmt.Println("連接成功")
}
