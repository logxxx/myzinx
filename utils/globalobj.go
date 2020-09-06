package utils

import (
	"000web/009zinx/ziface"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 409600,
	}

	GlobalObject.Reload()
}

func (g *GlobalObj) Reload() {
	wd, _ := os.Getwd()
	fmt.Println("wd:", wd)
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("Reload ReadFile err:", err)
		panic(err)
	}
	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		fmt.Println("Reload Unmarshal err:", err)
		panic(err)
	}
}
