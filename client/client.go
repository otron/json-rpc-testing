package main

import (
	"log"
	"github.com/ethereum/go-ethereum/rpc"
)

func getClient(printSMD bool) *rpc.Client {
	address := "http://localhost:9999"
	client, err := rpc.DialHTTP(address)
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	log.Printf("Connected to %s", address)
	if printSMD {
		smd := getSMD(client)
		log.Println(smd)
	}
	return client
}

func main() {
	client := getClient(true)
	var result interface{}
	err := client.Call(&result, "print.print", "what", "is this?")
	if err != nil {
		log.Fatalf("Error in message: %s", err)
	}
	log.Println("result:", result)
}

func getSMD(client *rpc.Client) interface{} {
	var result interface{}
	err := client.Call(&result, "GetSMD")
	if err != nil {
		log.Fatalf("Error retrieving SMD: %s", err)
	}
	return result
}
