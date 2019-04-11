package main

import (
	"net/http"
	"log"
	"flag"
	"os"
	"errors"
	"github.com/semrush/zenrpc"
	"github.com/semrush/zenrpc/smd"
)

type PrintService struct { zenrpc.Service }

type SMDService struct { zenrpc.Service }


func (ss SMDService) GetSMD() smd.Schema {
	return server.SMD()
}

func (ps PrintService) Print(a, b string) (int, error) {
	if a == "what" {
		return 1, errors.New("invalid string value because reasons")
	}
	log.Printf("'a'/param 1: '%s'. b/param 2: '%s'", a, b)
	return 0, nil
}

func (ps PrintService) PrintAny(value interface{}) (interface{}, error) {
	log.Printf("PrintAny: %s", value)
	return value, nil
}

var server zenrpc.Server
func main () {
	addr := flag.String("addr", "localhost:9999", "listen address")
	flag.Parse()

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	server = rpc
	rpc.Register("print", PrintService{})
	rpc.Register("", SMDService{})
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
