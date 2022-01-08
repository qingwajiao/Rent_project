package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server int

func (this *Server) Revd(arg int, resq *int) error {
	*resq = arg + 1
	return nil
}

func main() {

	c := new(Server)

	// 将Server对象注册到rpc中
	rpc.Register(c)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1235")
	if err != nil {
		fmt.Println("连接错误", err)
	}

	http.Serve(l, nil)

}
