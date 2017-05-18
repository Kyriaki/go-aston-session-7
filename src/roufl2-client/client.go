package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":6969")
	if err != nil {
		fmt.Println(err)
		return
	}

    cnonce := make([]byte, 0, 16) // cnonce buffer
    reader := make([]byte, 256)     // reading bufufer
 
    n, err := conn.Read(tmp)
    if err != nil {
        fmt.Println("read error:", err)
    }
        
    fmt.Println("got", n, "bytes.")
    cnonce = append(cnonce, tmp[:16]...)


	fmt.Println(string(buf))

}