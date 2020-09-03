package znet

import (
	"000web/009zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClosed  bool
	handleAPI ziface.HandleFunc
	ExitChan  chan bool
	Router    ziface.IRouter
}

func NewConnection(
	conn *net.TCPConn,
	connID uint32,
	router ziface.IRouter,
) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
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
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Printf("StartReader ReadFull err:%v\n", err)
			return
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Printf("StartReader Unpack err:%v\n", err)
			return
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Printf("StartReader ReadFull err:%v\n", err)
				return
			}
		}
		msg.SetData(data)

		req := &Request{
			conn: c,
			msg:  msg,
		}

		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(req)
	}
}

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("conn closed, cannot send msg.")
	}
	dp := NewDataPack()
	msg := NewMessage(msgId, data)
	binaryMsg, err := dp.Pack(msg)
	if err != nil {
		fmt.Printf("SendMsg Pack err:%v\n", err)
		return err
	}
	_, err = c.Conn.Write(binaryMsg)
	if err != nil {
		fmt.Printf("SendMsg Write err:%v\n", err)
		return err
	}
	return nil
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
