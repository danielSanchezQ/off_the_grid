package common

// Arguments for the Sleep command
// just taking the number of seconds to sleep the server call
type CommSleepArg struct {
	Seconds uint
}

// Arguments for the SimpleOperation command
// A and B are numbers to perform operation with
// Operators is a string representing the operation, should match any available operation
type CommSimpleOperationArg struct {
	A, B     int
	Operator string
}

// Base reply, message containing server computed data
type CommReply struct {
	Message string
}
