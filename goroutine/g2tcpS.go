package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		data := fmt.Sprintf("current time:%s\n", time.Now().Format("2006-01-02 15:04:05"))
		_, err := conn.Write([]byte(data))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 5)
	}

}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
