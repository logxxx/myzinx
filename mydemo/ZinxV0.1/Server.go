package main

import "000web/009zinx/znet"

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.Serve()
}
