package znet

import (
	"000web/009zinx/ziface"
	"fmt"
)

type MsgHandle struct {
	//存放每个msgId所对应的处理方法
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter, 0),
	}
}

func (h *MsgHandle) DoMsgHandler(req ziface.IRequest) {
	//1.从request中找到msgId
	handler, ok := h.Apis[req.GetMsgID()]
	if !ok {
		fmt.Printf("DoMsgHandler err:%v msgId:%v",
			"cannot match msgId to api", req.GetMsgID())
		return
	}

	//2.根据msgId调度对应的router业务
	handler.PreHandle(req)
	handler.Handle(req)
	handler.PostHandle(req)
}
func (h *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	//1.判断当前Msg绑定的API处理方法是否存在
	if _, ok := h.Apis[msgId]; ok {
		return
	}

	//2.添加msg与api的绑定关系
	h.Apis[msgId] = router
	fmt.Printf("AddRouter succ. msgId:%v\n", msgId)
}
