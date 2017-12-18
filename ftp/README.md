# ftp server

```golang
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port     = flag.String("port", "8000", "http port.")
	filepath = flag.String("filepath", ".", "ftp server dir.")
	logfile  = flag.String("logfile", "/var/log/ftp.log", "ftp log file.")
)

func logfileCreated() (*os.File, error) {
	fd, err := os.OpenFile(*logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return fd, nil

}

func main() {
	flag.Parse()

	fd, err := logfileCreated()
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	logger := log.New(fd, "INFO: ", log.Lshortfile)
	logger.SetFlags(logger.Flags() | log.LstdFlags)

	logger.Println("Start ftp server ...")
	listen_addr := fmt.Sprintf(":%s", *port)
	logger.Printf("Listen addr %s, ftp dir %s.\n", listen_addr, *filepath)
	logger.Fatal(http.ListenAndServe(listen_addr, http.FileServer(http.Dir(*filepath))))
}
```

## Usage

```golang
$ go build ftp.go
$ ./ftp -h

$ ./ftp
```
