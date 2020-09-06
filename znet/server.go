package znet

import (
	"000web/009zinx/utils"
	"000web/009zinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	MsgHandler ziface.IMsgHandle
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name:%v, listener HOST:%v PORT:%v",
		s.Name, s.IP, s.Port,
	)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion,
			fmt.Sprintf("%v:%v", s.IP, s.Port))
		if err != nil {
			fmt.Println("Start ResolveTCPAddr err", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ",
				s.IPVersion, "err:", err)
			return
		}
		fmt.Println("Start() ListenTCP() succ: ",
			"name=", s.Name)

		var cid uint32
		cid = 0

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Start AcceptTCP err:", err)
				continue
			}
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++
			dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("Server AddRouter Succ")
}
