package common

type CommSleepArg struct {
	Seconds uint
}

type CommSimpleOperationArg struct {
	A, B     int
	Operator string
}

type CommReply struct {
	Message string
}
