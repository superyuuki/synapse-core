package httpad

import (
	"math/rand"
	"synapse/synapse/logging"
	"synapse/synapse/pool"
	"time"
)

type HttpWork struct {
	lg logging.Logger
}

type HttpInput struct {
	target string
}

type HttpOutput struct {
	result string
}

func (h HttpWork) Process(input pool.Input) pool.Output {
	h.lg.Info("starting: " + input.(HttpInput).target)

	time.Sleep(time.Duration(int(time.Second) * rand.Intn(10)))

	str := input.(HttpInput).target + "wing"

	h.lg.Info("finished: " + str)

	return HttpOutput{result: str}
}
