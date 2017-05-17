package main

import (
	"fmt"
	"net"
)
func handleConnection(conn net.Conn) {

	conn.Write([]byte("Hello, world. \n"))
}

func main() {
	
	ln, err := net.Listen("tcp", ":6969")
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERROR")
		}
		go handleConnection(conn)
	}

}

