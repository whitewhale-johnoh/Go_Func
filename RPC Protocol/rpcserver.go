package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calc int // to register func in RPC Server

type Args struct { // argument
	A, B int
}

type Reply struct { // return value
	C int
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B //sum two value and put return value in struct
	return nil
}

func main() {
	rpc.Register(new(Calc))               // Create calc type instance and register in rpc server
	ln, err := net.Listen("tcp", ":6000") // get connection to tcp protocol 6000
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() // close wait for connection before main func close

	for {
		conn, err := ln.Accept() // return tcp connection when client connect
		if err != nil {
			continue
		}
		defer conn.Close() // close tcp session when main func close

		go rpc.ServerConn(conn)
	}
}
