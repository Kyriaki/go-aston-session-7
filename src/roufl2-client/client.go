package main

import (
	"fmt"
	"io"
	"net"
	"crypto/sha1"
	"encoding/hex"
)

func main() {

	ip := "127.0.0.1"
	username := "kyri"
	password := "go"


	conn, err := net.Dial("tcp", ":6969")
	if err != nil {
		fmt.Println(err)
		return
	}

    cn := make([]byte, 0, 16) // cnonce buffer
    reader := make([]byte, 256)     // reading bufufer
 
    n, err := conn.Read(reader)
    if err != nil {
        fmt.Println("read error:", err)
    }
    if n != 0{
    	cn = append(cn, reader[:16]...)	
    }

	cnonce := string(cn)
	resp :=computeResponse(ip, username, password, cnonce)
	conn.Write([]byte("AUTH " + username +" " + resp))
	fmt.Println("AUTH " + username + " " + resp)
	for{
		
	}
}

func computeResponse(ip string, username string, password string, cnonce string) string{
	h1 := sha1.New()
	response := sha1.New()
	io.WriteString(h1, ip + ":" + username + ":" + password)
	H1 := hex.EncodeToString(h1.Sum(nil))

	io.WriteString(response, H1 + ":" + cnonce)
	Response := hex.EncodeToString(response.Sum(nil))

	return Response
}