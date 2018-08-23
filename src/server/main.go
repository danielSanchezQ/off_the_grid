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
	log.Printf("Launching server app: serving on %s port: %s", NETWORK, PORT)
	//register available SimpleOperations
	operations := make(map[string]func(int, int) int)
	operations["+"] = func(i int, i2 int) int { return i + i2 }
	operations["-"] = func(i int, i2 int) int { return i - i2 }
	operations["*"] = func(i int, i2 int) int { return i * i2 }
	operations["/"] = func(i int, i2 int) int { return i / i2 }

	//initialize rpc handler
	simpleOperation := common.SipleOperation{
		AvailableOperations: operations,
	}
	server := rpc.NewServer()
	//register handler into server
	server.RegisterName("SimpleOperation", &simpleOperation)
	server.HandleHTTP("/", "/debug")
	listener, e := net.Listen(NETWORK, ":"+PORT)
	if e != nil {
		log.Fatal("Listening error:", e)
	} else {
		http.Serve(listener, nil)
	}
}
