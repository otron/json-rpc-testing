package main

import (
	"log"
	"github.com/ethereum/go-ethereum/rpc"
	net "net"
)

type CalculatorService struct {}

func (s *CalculatorService) Add(a, b int) (int, error) {
	log.Println("Add!", a, b)
	return a+b, nil
}


func main() {
	log.Println("Hello!");
	calculator := new(CalculatorService)
	server := rpc.NewServer()
	server.RegisterName("calculator", calculator)

	l, _ := net.ListenUnix("unix", &net.UnixAddr{Net: "unix", Name: "/tmp/calculator.sock"})
	for {
		c, _ := l.AcceptUnix()
		log.Println("data:", &c);
		codec := rpc.NewJSONCodec(c)
		log.Println(codec)
		go server.ServeCodec(codec, 0)
	}
}
