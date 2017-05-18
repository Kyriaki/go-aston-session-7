package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"math/rand"
	"crypto/sha1"
	"encoding/hex"
)
func handleConnection(conn net.Conn) {

	//conn.Write([]byte("Hello, world. \n"))
}

func main() {
	

	src := rand.NewSource(16374012946015784)

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
		fmt.Println("someone connected")
		
		conn.Write([]byte(generateCNonce(src)))
		go handleConnection(conn)
	}
	//computeResponse("127.0.0.1", "kyri", "aston", "0ab4f113b")

}

//func authClient()

func generateCNonce(src rand.Source) string {
	random := rand.New(src)
	randomNum := random.Int63()
	bs := strconv.FormatInt(randomNum, 16)
	return bs
}

func computeResponse(ip string, username string, password string, cnonce string) string{
	h1 := sha1.New()
	response := sha1.New()
	io.WriteString(h1, ip + ":" + username + ":" + password)
	H1 := hex.EncodeToString(h1.Sum(nil))
	fmt.Println(h1)

	io.WriteString(response, H1 + ":" + cnonce)
	Response := hex.EncodeToString(response.Sum(nil))

	return Response
}
