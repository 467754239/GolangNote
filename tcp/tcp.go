package main

import (
	"log"
	"net"
)

var RequestHeader = `golang tcp server.
`

func handleConn(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(RequestHeader))
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		log.Printf("new connection from %s\n", conn.RemoteAddr())
		if err != nil {
			continue
			log.Println(err)
		}
		go handleConn(conn)
	}
}
