package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Args rpc
type Args struct{}

//TimeServer rpc
type TimeServer int64

// GiveServerTime rpc
func (t *TimeServer) GiveServerTime(arguments *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}
func main() {
	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(listen, nil)
}
