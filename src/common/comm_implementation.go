package common

import (
	"fmt"
	"os"
	"time"
)

type SipleOperation struct {
	available_operations map[string]func(int, int) int
}

func (so *SipleOperation) Sleep(arg *CommSleepArg, reply *CommReply) error {
	time.Sleep(time.Duration(arg.Seconds) * time.Second)
	fmt.Fprintf(os.Stdout, "Finished waiting %d\n", arg.Seconds)
	reply.Message = fmt.Sprintf("Finished waiting %d", arg.Seconds)
	return nil
}

func (so *SipleOperation) SimpleOperation(arg *CommSimpleOperationArg, reply *CommReply) error {
	f, ok := so.available_operations[arg.Operator]
	if ok {
		res := f(arg.A, arg.B)
		fmt.Fprintf(os.Stdout, "Computed result (%d %s %d = %d)\n", arg.A, arg.Operator, arg.B, res)
		reply.Message = fmt.Sprintf("Computed result (%d %s %d = %d)", arg.A, arg.Operator, arg.B, res)
	} else {
		fmt.Fprintf(os.Stdout, "Could not match any operator %s\n", arg.Operator)
		reply.Message = fmt.Sprintf("Could not match any operator %s", arg.Operator)
	}
	return nil
}
