package main

import (
	"000web/009zinx/znet"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8998")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	for {
		//发送封包的message消息
		dp := znet.NewDataPack()
		msg := znet.NewMessage(0, []byte(
			"i'm client v0.5... "+
				"Today is "+time.Now().Format("2006-01-02T15:04:05")+
				""))
		packedMsg, err := dp.Pack(msg)
		if err != nil {
			fmt.Println("dp.Pack err:", err)
			return
		}
		_, err = conn.Write(packedMsg)
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}

		//先读取流中的head部分 得到id和dataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, binaryHead)
		if err != nil {
			fmt.Println("io.ReadFull err:", err)
			return
		}
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("dp.Unpack err:", err)
			return
		}

		//再根据dataLen进行二次读取，读出内容
		serverMsg := msgHead.(*znet.Message)
		serverMsg.Data = make([]byte, msgHead.GetMsgLen())
		_, err = io.ReadFull(conn, serverMsg.Data)
		if err != nil {
			fmt.Println("io.ReadFull err:", err)
			return
		}
		fmt.Println("recv server msg. Id:", serverMsg.Id,
			" len:", serverMsg.GetMsgLen(),
			" data:", string(serverMsg.GetData()))

		time.Sleep(1 * time.Second)
	}
}
