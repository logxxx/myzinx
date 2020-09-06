package main

import (
	"000web/009zinx/ziface"
	"000web/009zinx/znet"
	"fmt"
	"time"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	fmt.Println("recv from client: msgId=", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(
		0,
		[]byte("hehe hehe hehe"))
	if err != nil {
		fmt.Println("pingrouter handle sendmsg err:", err)
	}
}

type TimeRouter struct {
	znet.BaseRouter
}

func (p *TimeRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	fmt.Println("recv from client: msgId=", request.GetMsgID(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(
		1,
		[]byte(fmt.Sprintf("server time:%v",
			time.Now().Format("2006年01月02号 15点04分05秒"))))
	if err != nil {
		fmt.Println("TimeRouter handle sendmsg err:", err)
	}
}

func main() {
	s := znet.NewServer("[zinx V0.6]")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &TimeRouter{})
	s.Serve()
}
