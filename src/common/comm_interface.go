package common

// Interface of available rpc calls
// Sleep: sleep the conection gorutine for args.seconds seconds
// SimpleOperation: performs an operation in the form of arg.operator(arg.A, arg.B)
type ConnInterface interface {
	Sleep(arg *CommSleepArg, reply *CommReply) error
	SimpleOperation(arg *CommSimpleOperationArg, reply *CommReply) error
}
