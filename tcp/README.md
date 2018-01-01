## tcp

```golang
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
```




## tcp to http

```golang
package main

import (
	"log"
	"net"
)

var RequestHeader = `HTTP/1.1 200 OK
Content-Type: text/html
Last-Modified: 2017-12-21 14:25:00 
Server: Nginx 1.8

<html>
    <body>
	<h1 style="color:red">golang http server.</h1>
	</body>
</html>
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
```
