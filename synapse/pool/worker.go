package pool

import (
	"synapse/synapse/logging"
	"sync"
)

type Worker interface {
	Start()
}

type WorkerImpl struct {
	wg *sync.WaitGroup
	lg logging.Logger

	processor *Work

	input          chan Input
	output         chan Output
	globalShutdown chan Cancel
}

func NewWorker(wg *sync.WaitGroup, lg logging.Logger, processor *Work, input chan Input, output chan Output, globalShutdown chan Cancel) Worker {
	return WorkerImpl{
		wg:             wg,
		lg:             lg,
		processor:      processor,
		input:          input,
		output:         output,
		globalShutdown: globalShutdown,
	}
}

func (w WorkerImpl) Start() {
	defer w.wg.Done()

	go func() {
		for true {
			select {
			case inputObject := <-w.input:
				pr := *w.processor //dereference

				w.output <- pr(inputObject)
			case <-w.globalShutdown:
				return //commit suicide
			}
		}
	}()
}
