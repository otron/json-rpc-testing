package services

import (
	"github.com/semrush/zenrpc"
	"github.com/semrush/zenrpc/smd"
	"errors"
	"log"
)


type PrintService struct { zenrpc.Service }

type SMDService struct {
	zenrpc.Service
	Server *zenrpc.Server
}

func (ss SMDService) GetSMD() smd.Schema {
	return ss.Server.SMD()
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
