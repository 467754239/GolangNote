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
	addr   string
	err    error
	fd     *os.File
	logger *log.Logger

	port     = flag.String("port", "8000", "http port.")
	filepath = flag.String("filepath", ".", "ftp server dir.")
	enable   = flag.Bool("enable", false, "enable stdout log to file.")
	logfile  = flag.String("logfile", "/var/log/ftp.log", "ftp log file.")
)

func init() {
	flag.Parse()
	addr = fmt.Sprintf(":%s", *port)

	fd, err = os.OpenFile(*logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	//defer fd.CLose()

	logger = log.New(fd, "INFO: ", log.Lshortfile)
	logger.SetFlags(logger.Flags() | log.LstdFlags)
}

func main() {
	if *enable {
		logger.Println("Start ftp server ...")
		logger.Printf("Listen addr %s, ftp dir %s.\n", addr, *filepath)
	} else {
		fmt.Println("Start ftp server ...")
		fmt.Printf("Listen addr %s, ftp dir %s.\n", addr, *filepath)
	}

	logger.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir(*filepath))))
}
```

## Usage

```golang
$ go build ftp.go
$ ./ftp -h

$ ./ftp
```
