package socketApi

import (
	"net"

	log "github.com/auxten/logrus"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		// read from the connection
		var buf = make([]byte, 1024)
		log.Debug("start to read from conn")
		n, err := conn.Read(&buf)
		if err != nil {
			log.Error("conn read error:", err)
			return
		}
		log.Debug("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func SocketServerStart() {
	lis, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatal("Listen error:", err)
		return
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}

			log.Fatal("Accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		log.Debug("accept a new connection")
		go handleConn(conn)
	}
}
