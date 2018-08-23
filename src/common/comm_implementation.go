package common

import (
	"fmt"
	"log"
	"time"
)

// Base struct implementation for ConnInterface
type SipleOperation struct {
	AvailableOperations map[string]func(int, int) int
}

//Sleep the server connection gorutine for arg.seconds seconds
func (so *SipleOperation) Sleep(arg *CommSleepArg, reply *CommReply) error {
	time.Sleep(time.Duration(arg.Seconds) * time.Second)
	log.Printf("Finished waiting %d\n", arg.Seconds)
	reply.Message = fmt.Sprintf("Finished waiting %d", arg.Seconds)
	return nil
}

// Execute an already registered operation
// If revieved operation (arg.Operator) is not found just reply informing about it
func (so *SipleOperation) SimpleOperation(arg *CommSimpleOperationArg, reply *CommReply) error {
	f, ok := so.AvailableOperations[arg.Operator]
	if ok {
		// autorecover gorutine from any panic in the f(arg.A, arg.B) call
		defer func() {
			recoverResult := recover()
			if recoverResult != nil {
				log.Printf("An error occured while performing operation %s\n", recoverResult)
				reply.Message = fmt.Sprintf("An error occured while performing operation %s\n", recoverResult)
			}
		}()
		res := f(arg.A, arg.B) // excute operation
		log.Printf("Computed result (%d %s %d = %d)\n", arg.A, arg.Operator, arg.B, res)
		reply.Message = fmt.Sprintf("Computed result (%d %s %d = %d)", arg.A, arg.Operator, arg.B, res)
	} else {
		log.Printf("Could not match any operator %s\n", arg.Operator)
		reply.Message = fmt.Sprintf("Could not match any operator %s", arg.Operator)
	}
	return nil
}
