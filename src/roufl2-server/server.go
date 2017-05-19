package main

import (
	"fmt"
	"io"
	"bufio"
	"net"
	"time"
	"strings"
	"strconv"
	"math/rand"
	"crypto/sha1"
	"encoding/hex"
)

func main() {

	cltPsw := "go"
	src := rand.NewSource(16374012946015784)

	ln, err := net.Listen("tcp", ":6969")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("someone connected")
		buf := bufio.NewReader(conn)

		cnonce := generateCNonce(src)
		conn.Write([]byte(cnonce))	

		for x := range time.Tick(50*time.Millisecond){
		 	_ = x
			n, _ := buf.ReadString('\n')
			if n != ""{
				fmt.Println(handleQueries(n, conn, cnonce, cltPsw))	
			}
		}
	}
}

func authClient(auth string, cnonce string, clientPassword string) int {
	params := strings.Split(auth, " ")
	username := params[1]
	respClient := params[2]

	respServer := computeResponse("127.0.0.1", username, clientPassword, cnonce)
	fmt.Println("Response server: " + respServer + "\nResponse client: " + respClient)
	if respServer == respClient{
		return 1
	} else {
		return 0
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

func serverAnswer(code int, data string) string{
	switch code{
		case 69:
			return "69 OK " + data 
		case 51:
			return "51 NOK"
		case 333:
			return "333 BADASS"
		case 444:
			return "444 BADAUTH"
		case 555:
			return "555 BADARGUMENT"
		case 666:
			return "666 INTERNALERROR"
		case 777:
			return "777 BADREQUEST"
		case 888:
			return "888 BADFORMAT"
		case 999:
			return "999 BADLANGUAGE"
		default:
			return "666 INTERNALERROR"
	}
}

func handleQueries(query string, conn net.Conn, cnonce string, clientPass string) string {

	queryArgs := strings.Split(strings.TrimSuffix(query, "\n"), " ")
	if len(queryArgs) != 6{
		return serverAnswer(777, "")
	} else {
		command := queryArgs[0]
		language := queryArgs[1]
		data := queryArgs[2]
		output := queryArgs[3]
		clientResponse := queryArgs[4]
		username := queryArgs[5]

		authQuery := "AUTH "+ username + " "+ clientResponse
		if checkCommand(command) == false {
			return serverAnswer(333, "")
		}
		if checkLanguages(language) == false {
			return serverAnswer(999, "")
		}
		if checkOutput(output) == false {
			return serverAnswer(555, "")
		}
		if authClient(authQuery, cnonce, clientPass) == 0{
			return serverAnswer(444, "")
		}
		return serverAnswer(69, data)
	}
}

func checkCommand(command string) bool{
	if command != "SERIALIZE" && command != "UNSERIALIZE" {
		return false
	} else {
		return true
	}
}

func checkLanguages(language string) bool{

	languages := make(map[int]string)
	languages[0] = "Go"
	for lng := range languages{
		if language == languages[lng]{
			return true
		} 
	}
	return false
}

func checkOutput(output string) bool{
	if output != "JSON" && output != "XML" && output != "Binary"{
		return false
	} else {
		return true
	}
}