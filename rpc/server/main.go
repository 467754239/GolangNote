package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/467754239/GolangNote/rpc/common"
)

type MathService struct {
}

func (m *MathService) Add(request *common.AddRequest, reply *common.AddResponse) error {
	log.Printf("call add:%v", request)
	reply.Result = request.M + request.N
	return nil
}

func main() {
	matchService := new(MathService)

	rpc.Register(matchService)

	ln, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}
	rpc.Accept(ln)

}
