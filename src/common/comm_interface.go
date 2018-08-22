package common

type ConnInterface interface {
	Sleep(arg *CommSleepArg, reply *CommReply) error
	SimpleOperation(arg *CommSimpleOperationArg, reply *CommReply) error
}
