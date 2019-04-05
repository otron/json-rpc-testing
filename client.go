package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	a, b int
}

func main() {
	client, err := rpc.Dial("unix", "/tmp/calculator.sock")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	args := &Args{a: 2, b: 3}
	var result int
	client.Call("calculator", args, &result)
	log.Println("result:", result)
}
