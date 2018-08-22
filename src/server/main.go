package main

import (
	"common"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var NETWORK = "tcp"
var PORT = "8080"

func main() {
	log.Print("Launching server app: serving on " + NETWORK + " " + PORT)
	operations := make(map[string]func(int, int) int)
	operations["+"] = func(i int, i2 int) int { return i + i2 }
	operations["-"] = func(i int, i2 int) int { return i - i2 }
	operations["*"] = func(i int, i2 int) int { return i * i2 }
	operations["/"] = func(i int, i2 int) int { return i / i2 }

	simpleOperation := common.SipleOperation{
		AvailableOperations: operations,
	}
	server := rpc.NewServer()
	server.RegisterName("SimpleOperation", &simpleOperation)
	server.HandleHTTP("/", "/debug")
	listener, e := net.Listen(NETWORK, ":"+PORT)
	if e != nil {
		log.Fatal("Listening error:", e)
	}
	http.Serve(listener, nil)
}
