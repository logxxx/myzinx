package znet

import (
	"000web/009zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	handleAPI ziface.HandleFunc
	ExitChan chan bool
}

func NewConnection(
	conn *net.TCPConn,
	connID uint32,
	callbckApi ziface.HandleFunc,
	) (*Connection) {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: callbckApi,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("Connetion.Stop().connId=",
		c.ConnID, " remote addr=", c.Conn.RemoteAddr())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err:", err)
			continue
		}
		fmt.Println("recv buf succ:cnt=", cnt)

		err = c.handleAPI(c.Conn, buf, cnt)
		if err != nil {
			fmt.Println("handleAPI err:", err,
				" connId:", c.ConnID)
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Connection Start... ConnID=", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID=", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Stop conn.Close() err:", err)
		return
	}

	close(c.ExitChan)

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}