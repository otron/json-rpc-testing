package main

import (
	"net/http"
	"log"
	"flag"
	"os"
	"errors"
	"github.com/semrush/zenrpc"
)

type PrintService struct { zenrpc.Service }

type ArithService struct { zenrpc.Service }

type Quotient struct {
	Quo, Rem int
}

func (as ArithService) Divide(a, b int) (quo *Quotient, err error) {
	if b == 0 {
		return nil, errors.New("divide by zero")
	} else if b == 1 {
		return nil, zenrpc.NewError(401, errors.New("we do not serve 1"))
	}

	return &Quotient{
		Quo: a / b,
		Rem: a % b,
	}, nil
}

func (ps PrintService) Print(a, b string) int {
	log.Println("here are yo strings: '%s' but also '%s'", a, b)
	return 0
}

func main () {
	addr := flag.String("addr", "localhost:9999", "listen address")
	flag.Parse()

	rpc := zenrpc.NewServer(zenrpc.Options{ExposeSMD: true})
	rpc.Register("print", PrintService{})
	// rpc.Register("arith", testdata.ArithService{})
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting printserver on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
