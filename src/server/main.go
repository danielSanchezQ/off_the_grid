package main

import (
	"common"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	simpleOperation := new(common.SipleOperation)
	server := rpc.NewServer()
	server.RegisterName("SimpleOperation", simpleOperation)
	server.HandleHTTP("/", "/debug")
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("Listening error:", e)
	}
	http.Serve(l, nil)
}
