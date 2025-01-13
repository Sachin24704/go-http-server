package main

import (
	"fmt"
	"log"
	"net"
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
	go func () {
		for {
			c, err := ll.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleConnection(c)
		}
	} ()
}

func handleConnection(c net.Conn) {
	defer c.Close()
	// read http formated request
	
	c.Read()
}
