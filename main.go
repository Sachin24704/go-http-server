package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// listen by tcp connection on 8080
	ll, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ll.Close()
	fmt.Println("server listening on port 8080")
	// establish a connection with the client
	for {
		c, err := ll.Accept()
		if err != nil {
			log.Println(err)
				continue
		}
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// read http formated request
	reader := bufio.NewReader(conn)
	reqLine, err := reader.ReadString('\n')
	if err != nil {
		log.Panicln(err)
		// as we cannot read request.
		return 
	}
	// parse the request
	fmt.Println("request:", reqLine)
	parseRequest(reqLine, conn)
}

func parseRequest(req string, conn net.Conn) {
	reqSlice := strings.Fields(req)
	if len(reqSlice) < 3 {
		sendResponse(conn, "Invalid Request", "400 Bad Request")
		return
	}
	path := reqSlice[1]
	switch reqSlice[0] {
	case "GET" : {
		if path == "/" {
			sendResponse(conn, "200 OK", "Welcome to the Simple HTTP Server in go !!!")
		} else {
			sendResponse(conn, "404 Not Found", "Page Not Found")
		}
	}
	default :
		sendResponse(conn, "Invalid Request", "400 Bad Request")
	}
}

func sendResponse(conn net.Conn, status string, body string) {
	// basic http response format
	response := fmt.Sprintf(
		"HTTP/1.1 %s\r\n"+
			"Content-Length: %d\r\n"+
			"Content-Type: text/plain\r\n"+
			"Connection: close\r\n"+
			"\r\n"+
			"%s",
		status, len(body), body,
	)
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Println(err)
	}
}


