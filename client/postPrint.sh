#!/bin/bash
curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"print.print","params": ["yow", "param2"], "id": "1"}' http://localhost:9999
#curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"print.print","params": {"a": "butts", "b": "yow bby"}, "id": "1"}' http://localhost:9999
#curl -X POST -H "Content-Type: application/json" --data ' {"jsonrpc": "2.0", "method": "print.print", "params": {"a": "333", "b": "91919"}, "id": 3}' http://localhost:9999
