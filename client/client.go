package main

import (
	"log"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.DialHTTP("http://localhost:9999/print")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	getSMD(client)
	var result interface{}
	err = client.Call(&result, "print", "what", "is this?")
	if err != nil {
		log.Fatalf("Error in message: %s", err)
	}
	log.Println("result:", result)
}

func getSMD(client *rpc.Client) {
	var result interface{}
	err := client.Call(&result, "GetSMD")
	if err != nil {
		log.Fatalf("oh boy, couldn't get the SMD stuff like that it seems. %s", err)
	} else {
		log.Println("SMD: %s", result)
	}
}
