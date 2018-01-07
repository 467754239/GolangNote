package audit

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/467754239/GolangNote/lh/config"
)

var waitgroup sync.WaitGroup

type Sender struct {
	addr string
}

func NewSender(addr string) *Sender {
	return &Sender{
		addr: addr,
	}
}

func (this *Sender) connect() (net.Conn, error) {
	conn, err := net.Dial("tcp", this.addr)
	return conn, err
}

func (this *Sender) send() {
	conn, err := this.connect()
	if err != nil {
		os.Exit(1)
		log.Print(err)
	}
	defer conn.Close()

	var system_env config.SystemEnv
	system_env.Setenv()
	body, err := json.Marshal(system_env)
	if err != nil {
		log.Print(err)
	}
	_, err = fmt.Fprintf(conn, "%v\n", string(body))
	if err != nil {
		log.Print(err)
	}
	waitgroup.Done()
}

func (this *Sender) Start() {
	waitgroup.Add(1)
	go this.send()
	waitgroup.Wait()
}
