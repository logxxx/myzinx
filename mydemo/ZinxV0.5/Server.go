package main

import (
	"000web/009zinx/ziface"
	"000web/009zinx/znet"
	"fmt"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	fmt.Println("recv from client: msgId=", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(
		1,
		[]byte("hehe hehe hehe"))
	if err != nil {
		fmt.Println("pingrouter handle sendmsg err:", err)
	}
}

func main() {
	s := znet.NewServer("[zinx V0.5]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
