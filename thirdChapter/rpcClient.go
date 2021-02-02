package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Args rpc
type Args struct{}

func main() {
	args := Args{}
	var reply int64
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("client: ", err)
	}
	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("call error: ", err)
	}
	fmt.Printf("Server Time: %d\n", reply)
}
