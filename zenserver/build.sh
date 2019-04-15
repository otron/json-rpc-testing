#!/bin/bash

serverFile=zenserver.go
servicesFile=services/services.go
zenrpcOutServices=services_zenrpc.go

zenrpc $servicesFile
go build $serverFile
