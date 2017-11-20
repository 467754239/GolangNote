package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/467754239/GolangNote/monitor/common"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	// 构造sender
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric),
	}

}

func (s *Sender) Start() {
	/*
		建立连接
		循环从s.ch里面读取metric
		序列化metric
		发送数据
	*/

	go func() {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		// version 1
		/*
			for {
				metric := <-ch
				buf, _ := json.Marshal(metric)
				fmt.Fprintf(conn, "%s\n", metric)
			}
		*/
		// version 2
		for metric := range s.ch {
			buf, _ := json.Marshal(metric)
			fmt.Fprintf(conn, "%s\n", buf)
		}
	}()
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
