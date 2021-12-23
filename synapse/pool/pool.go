package pool

import (
	"synapse/synapse/logging"
	"sync"
)

type Pool interface {
	// Populate
	//
	//	Outputs two channels: one that is fired when everything is done and one that fires for every completed job
	//
	Populate(supplier func() Input) (chan Cancel, chan Output)
}

type PoolImpl struct {
	lg   logging.Logger
	work *Work

	maxTasks   int
	maxWorkers int

	globalShutdown chan Cancel
}

func NewPool(globalShut chan Cancel, lg logging.Logger, work *Work, maxTasks int, maxWorkers int) Pool {
	return PoolImpl{
		globalShutdown: globalShut,
		lg:             lg,
		work:           work,
		maxTasks:       maxTasks,
		maxWorkers:     maxWorkers,
	}
}

func (p PoolImpl) Populate(supplier func() Input) (chan Cancel, chan Output) {

	input := make(chan Input, p.maxTasks)
	output := make(chan Output, p.maxTasks)
	completion := make(chan Cancel)
	wg := &sync.WaitGroup{}

	//make workers
	for i := 0; i < p.maxWorkers; i++ {
		wg.Add(1)

		NewWorker(wg, p.lg, p.work, input, output, p.globalShutdown).Start()
	}

	//populate input channel
	for i := 0; i < p.maxTasks; i++ {
		input <- supplier()
	}

	go func() {
		close(input)

		// Wait for all goroutines to finish processing
		wg.Wait() //should be reduced if termination signal is sent anyways

		// Close output channel since all workers have finished processing
		close(output)

		completion <- StopSuccess
	}()

	return completion, output

}
