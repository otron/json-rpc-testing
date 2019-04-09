#!/bin/bash

serverFile=zenserver.go
zenrpcOut=main_zenrpc.go
buildOut=zenserver

if [ -e $buildOut ]; then 
    rm $buildOut
fi
if [ -e $zenrpcOut ]; then
    rm $zenrpcOut
fi
zenrpc $serverFile
go build $serverFile $zenrpcOut
