# running/building the zenserver.go example

1. install the zenrpc generator: `go get github.com/semrush/zenrpc/zenrpc`
2. run `zenrpc zenserver.go`. This will create a file `main_zenrpc.go`
3. build: `go build zenserver.go main_zenrpc.go`
4. run the `zenserver` binary
