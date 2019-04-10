# running/building the zenserver.go example

1. install the zenrpc generator: `go get github.com/semrush/zenrpc/zenrpc`
2. run `zenrpc zenserver.go`. This will create a file `main_zenrpc.go`
3. build: `go build zenserver.go main_zenrpc.go`
4. run the `zenserver` binary


## Development setup

### 1. Install Go

Find it with your distro's package manager or at https://golang.org/dl/

### 2. Install dependencies with Go

Run `go install`


### Debugging

https://github.com/go-delve/delve

## Building

1. generate zenrpc-stuff with the zenrpc binary:
  - `./../../bin/zenrpc Server.go`
2. build the binaries with `go build`
  - `go build Server.go main_zenrpc.go`
