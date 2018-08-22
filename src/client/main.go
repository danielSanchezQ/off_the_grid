package main

import (
	"common"
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

type ConnClient struct {
	Client         *rpc.Client
	RecieveChannel chan common.CommReply
	WaitingGroup   sync.WaitGroup
}

func (cc *ConnClient) callSleep(seconds uint) {
	cc.WaitingGroup.Add(1)
	go func() {
		defer cc.WaitingGroup.Done()
		args := &common.CommSleepArg{Seconds: seconds}
		var reply common.CommReply
		err := cc.Client.Call("SimpleOperation.Sleep", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		} else {
			cc.RecieveChannel <- reply
		}
	}()
}

func (cc *ConnClient) callSimpleOperation(a int, b int, operator string) {
	cc.WaitingGroup.Add(1)
	go func() {
		defer cc.WaitingGroup.Done()
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
			cc.RecieveChannel <- reply
		}
	}()
}
func (cc *ConnClient) asArrive() {
	go func() {
		for reply := range cc.RecieveChannel {
			fmt.Println(reply.Message)
		}
	}()
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	} else {
		connclient := &ConnClient{
			Client:         client,
			RecieveChannel: make(chan common.CommReply),
			WaitingGroup:   sync.WaitGroup{}}

		connclient.asArrive()
		for i := 10; i >= 0; i-- {
			connclient.callSleep(uint(i))
			for j := 0; j <= 10; j++ {
				for _, op := range [4]string{"+", "-", "*", "/"} {
					connclient.callSimpleOperation(i, j, op)
				}
			}
		}
		connclient.WaitingGroup.Wait()
	}
}
