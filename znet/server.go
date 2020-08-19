package znet

import (
	"000web/009zinx/ziface"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port int
}

func NewServer(name string) ziface.IServer {
	s := &Server {
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 8999,
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("[Server]Start.ip=%v port=%v", s.IP, s.Port)
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

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Start AcceptTCP err:", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("start Read err:", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("start Write err:", err)
						continue
					}
					fmt.Println("[", time.Now().Format("2006/01/02 15:04:05"), "]",
						"Receive from client:", conn.LocalAddr(),
						" content:", string(buf))
				}
			}()
		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	select{}
}
