package main

import (
	"fmt"
	"net/rpc"
)

var resp int

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1235")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}

	err = client.Call("Server.Revd", 66, &resp)

	if err != nil {
		fmt.Println("rpc调用失败", err)
		return
	}

	fmt.Println("rpc调用返回的结果为：", resp)

}
