package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"strconv"
	"math/rand"
	"crypto/sha1"
	"encoding/hex"
)
func handleConnection(conn net.Conn) {
	auth := make([]byte, 0, 1024) // authRequest buffer
    reader := make([]byte, 256)     // reading bufufer
 
    n, err := conn.Read(reader)
    if err != nil {
        fmt.Println("read error:", err)
    }
    if n != 0{
    	auth = append(auth, reader[:n]...)	
    }
    fmt.Println(string(auth))
}

func main() {
	
	cltPsw := "go"
	src := rand.NewSource(16374012946015784)

	ln, err := net.Listen("tcp", ":6969")
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("someone connected")
		
		cnonce := generateCNonce(src)
		conn.Write([]byte(cnonce))

		auth := make([]byte, 0, 1024) // authRequest buffer
	    reader := make([]byte, 256)     // reading bufufer
	 
	    n, err := conn.Read(reader)
	    if err != nil {
	        fmt.Println("read error:", err)
	    }

	    if n != 0{
	    	auth = append(auth, reader[:n]...)
	    }
	    serverResponse := authClient(string(auth), cnonce, cltPsw)
		fmt.Println(serverResponse)
	}
}

func authClient(auth string, cnonce string, clientPassword string) int {
	params := strings.Split(auth, " ")
	username := params[1]
	respClient := params[2]

	respServer := computeResponse("127.0.0.1", username, clientPassword, cnonce)
	fmt.Println("Response server: " + respServer + "\nResponse client: " + respClient)
	if respServer == respClient{
		return 200
	} else {
		return 888
	}

}

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
	
	io.WriteString(response, H1 + ":" + cnonce)
	Response := hex.EncodeToString(response.Sum(nil))

	return Response
}
