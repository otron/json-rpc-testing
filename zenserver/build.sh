#!/bin/bash

serverFile=zenserver.go
servicesFile=services/services.go
zenrpcOutServices=services_zenrpc.go
zenrpcOutMain=main_zenrpc.go

zenrpc $serverFile
zenrpc $servicesFile
go build $serverFile $zenrpcOutMain
