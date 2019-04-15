package main

import (
	"net/http"
	"log"
	"flag"
	"os"
	"github.com/semrush/zenrpc"
	"github.com/otron/json-rpc-testing/zenserver/services"
)


func main () {
	addr := flag.String("addr", "localhost:9999", "listen address")
	flag.Parse()

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	smdService := services.SMDService{Server: &rpc}

	rpc.Register("print", services.PrintService{})
	rpc.Register("", smdService)
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
