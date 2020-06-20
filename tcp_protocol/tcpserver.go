package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096) //  create 4096 byte slice

	for {
		n, err := c.Read(data) //  read the date received from client
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // print out data

		_, err = c.Write(data[:n]) // send data to client
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000") // receive connection through 8000 in TCP Protocol
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() // close the connection before main function over

	for {
		conn, err := ln.Accept() // return the connection when the client's connected
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()      // close the tcp session before the main func finished
		go requestHandler(conn) // run the func that dealt the packet through go routine
	}
}
