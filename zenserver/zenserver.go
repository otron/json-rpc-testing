package main

import (
	"net/http"
	"log"
	"flag"
	"os"
	"github.com/semrush/zenrpc"
	"github.com/semrush/zenrpc/smd"
	"github.com/otron/json-rpc-testing/zenserver/services"
)

type SMDService struct { zenrpc.Service }

func (ss SMDService) GetSMD() smd.Schema {
	return server.SMD()
}


var server zenrpc.Server
func main () {
	addr := flag.String("addr", "localhost:9999", "listen address")
	flag.Parse()

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	server = rpc

	rpc.Register("print", services.PrintService{})
	rpc.Register("", SMDService{})
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
