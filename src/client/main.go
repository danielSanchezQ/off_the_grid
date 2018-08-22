package main

import (
	"common"
	"fmt"
	"log"
	"net/rpc"
)

type ConnClient struct {
	client *rpc.Client
}

func (cc *ConnClient) callSleep(seconds uint) common.CommReply {
	args := &common.CommSleepArg{Seconds: seconds}
	var reply common.CommReply
	err := cc.client.Call("SimpleOperation.Sleep", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	connclient := &ConnClient{client: client}

	fmt.Println(connclient.callSleep(10))
}
