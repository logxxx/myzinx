package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")

	time.Sleep(1*time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err !=nil {
		fmt.Println("Dial err:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello world~"))
		if err != nil {
			fmt.Println("write conn err:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("client receive from server: len=%v", cnt)

		time.Sleep(1*time.Second)
	}
}