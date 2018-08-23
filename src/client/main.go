package main

import (
	"common"
	"log"
	"net/rpc"
	"sync"
)

// RPC client implementation
// Owns a channel for registering connection results
// Controls spawned gorutines through a Waiting group
type ConnClient struct {
	Client         *rpc.Client           //rpc client
	ReceiveChannel chan common.CommReply //channel where connection results will be dumped
	WaitingGroup   sync.WaitGroup        // spawned gorutines synchronization group
}

// Execute a Sleep call to remote server
// Spawns a new goruting that will store the result into the receive channel
func (cc *ConnClient) callSleep(seconds uint) {
	cc.WaitingGroup.Add(1) // increment spawned gorutines to wait
	go func() {
		defer cc.WaitingGroup.Done() // decrement spawned gorutines to wait
		args := &common.CommSleepArg{Seconds: seconds}
		var reply common.CommReply
		err := cc.Client.Call("SimpleOperation.Sleep", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		} else {
			cc.ReceiveChannel <- reply
		}
	}()
}

// Execute a SimpleOperation call to remote server
// Spawns a new goruting that will store the result into the receive channel
func (cc *ConnClient) callSimpleOperation(a int, b int, operator string) {
	cc.WaitingGroup.Add(1) // increment spawned gorutines to wait
	go func() {
		defer cc.WaitingGroup.Done() // decrement spawned gorutines to wait
		args := &common.CommSimpleOperationArg{
			A:        a,
			B:        b,
			Operator: operator,
		}
		var reply common.CommReply
		err := cc.Client.Call("SimpleOperation.SimpleOperation", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		} else {
			cc.ReceiveChannel <- reply
		}
	}()
}

//Spawn a gorutine that reads from the receive channel and print its content
func (cc *ConnClient) asArrive() {
	go func() {
		for reply := range cc.ReceiveChannel {
			log.Println(reply.Message)
		}
	}()
}

func main() {
	//connect to server
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error while connecting to server:", err)
	} else {
		//instantiate the connection manager
		connclient := &ConnClient{
			Client:         client,
			ReceiveChannel: make(chan common.CommReply),
			WaitingGroup:   sync.WaitGroup{}}

		//spawn reading gorutine
		connclient.asArrive()
		//execute operations
		for i := 10; i >= 0; i-- {
			connclient.callSleep(uint(i)) //calling sleep in reverse order to check gorutines do not wait to each other
			for j := 0; j <= 10; j++ {
				for _, op := range [4]string{"+", "-", "*", "/"} {
					//calling operations, including division by 0, server should autorecover
					connclient.callSimpleOperation(i, j, op)
				}
			}
		}
		//wait for operations to finish
		connclient.WaitingGroup.Wait()
	}
}
