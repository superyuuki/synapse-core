package pool

type Cancel int

const (
	StopSuccess = iota
	StopTerminated
	StopFailed
	StopContextual
)
