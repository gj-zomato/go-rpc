package main

import (
	mathFunc "go-rpc/math"
	"log"
	"net"
	"net/rpc"
)

func main() {
	arith := new(mathFunc.Arith)
	rpc.RegisterName("Arith", arith)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	log.Println("RPC server listening on port 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
