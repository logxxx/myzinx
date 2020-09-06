package ziface

type IMsgHandle interface {
	DoMsgHandler(req IRequest)
	AddRouter(msgId uint32, router IRouter)
}
