package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

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
		ch:   make(chan *common.Metric, 1),
	}

}

func (s *Sender) connect() net.Conn {
	n := 1000 * time.Millisecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Warn(err)
			time.Sleep(n)
			n = n * 2
			if n > time.Second*30 {
				n = time.Second * 30
			}
			continue
		}
		return conn
	}

}

func (s *Sender) Start() {
	/*
		建立连接
		循环从s.ch里面读取metric
		序列化metric
		发送数据
	*/

	/*
		for {
			metric := <-ch
			buf, _ := json.Marshal(metric)
			fmt.Fprintf(conn, "%s\n", metric)
		}
	*/

	// 定时定量发送
	conn := s.connect()

	// 蓄水池 4k
	w := bufio.NewWriter(conn)
	log.Infof("local addr:%v\n", conn.LocalAddr())
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case metric := <-s.ch:
			buf, _ := json.Marshal(metric)
			_, err := fmt.Fprintf(w, "%s\n", buf)
			if err != nil {
				conn.Close()
				conn = s.connect()
				w = bufio.NewWriter(conn)
				log.Infof("local addr:%v\n", conn.LocalAddr())
			}
		case <-ticker.C:
			err := w.Flush()
			if err != nil {
				conn.Close()
				conn = s.connect()
				w = bufio.NewWriter(conn)
				log.Infof("local addr:%v\n", conn.LocalAddr())
			}
		}
	}
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
