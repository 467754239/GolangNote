package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port     = flag.String("port", "8000", "http port.")
	filepath = flag.String("filepath", ".", "ftp server dir.")
)

func main() {
	flag.Parse()
	fmt.Println("Start ftp server ...")
	listen_addr := fmt.Sprintf(":%s", *port)
	fmt.Printf("Listen addr %s, ftp dir %s ...\n", listen_addr, *filepath)
	log.Fatal(http.ListenAndServe(listen_addr, http.FileServer(http.Dir(*filepath))))
}
