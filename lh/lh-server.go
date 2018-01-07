package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/467754239/GolangNote/lh/common"
	"github.com/467754239/GolangNote/lh/config"
	"github.com/467754239/GolangNote/lh/models"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(buf[:n]))

	db, err := models.NewConnect()
	if err != nil {
		log.Fatal(err)
	}

	var user_audit common.UserAudit
	err = json.Unmarshal(buf[:n], &user_audit)
	if err != nil {
		fmt.Println("error:", err)
	}
	affected, err := db.Insert(user_audit)
	if affected == 1 {
		log.Print("Insert sucess.")
	} else {
		log.Printf("Insert failed, info:%v\n", err)
	}

}

func main() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.VKLH_AUDIT_PROD_PORT))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}
