package pool

import (
	"context"
	"github.com/mum4k/termdash/cell"
	"synapse/synapse/logging"
)

type Start interface {
	SpawnInput() Input
}

type Holder interface {
	Init() (chan Start, chan Cancel)
}

type HolderImpl struct {
	ctx context.Context
	lg  logging.Logger

	work *Work

	maxTasks   int
	maxWorkers int
}

func NewHolder(ctx context.Context, lg logging.Logger, work *Work, maxTasks int, maxWorkers int) Holder {
	return HolderImpl{
		ctx:        ctx,
		lg:         lg,
		work:       work,
		maxTasks:   maxTasks,
		maxWorkers: maxWorkers,
	}
}

func (h HolderImpl) Init() (chan Start, chan Cancel) {
	start := make(chan Start) //big start
	stop := make(chan Cancel) //big stop

	go func() {
		isRunning := false

		for true {
			select {
			case <-h.ctx.Done():
				stop <- StopContextual

			case cancelSig := <-stop:

				if isRunning {
					isRunning = false

					switch cancelSig {
					case StopContextual:
					case StopSuccess:
					case StopFailed:
					case StopTerminated:

					}

				} else {
					h.lg.Error("Cannot stop http: http not running")
				}

			case startSig := <-start:

				if !isRunning {
					isRunning = true

					finout, logout := NewPool(stop, h.lg, h.work, h.maxTasks, h.maxWorkers).Populate(func() Input {
						return startSig.SpawnInput()
					})

					//TODO logging

					go func() {

						for range logout {
							h.lg.Prefixed("completed", "task", cell.ColorLime)
						}
						//log output

					}()

					go func() {
						out := <-finout
						stop <- out

						//if finished, fire complete signal
					}()
				}

			}
		}

	}()

	return start, stop
}
