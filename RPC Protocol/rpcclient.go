package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:6000") //Connect to RPC Server
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() //close RPC

	args := &Args{1, 2}
	reply := new(Reply)
	err = client.Call("Calc.Sum", args, reply) // call Calc.Sum func
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.C)

	args.A = 4
	args.B = 9
	sumCall := client.Go("Calc.Sum", args, reply, nil) //call goroutine Calc.Sum func
	<-sumCall.Done
	fmt.Println(reply.C)
}
