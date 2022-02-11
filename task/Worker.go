package task

// Worker type holds the job channel and passed worker threadpool
type Worker struct {
	jobChannel  chan Runnable
	workerPool  chan chan Runnable
	closeHandle chan bool
}

// NewWorker creates the new worker
func NewWorker(workerPool chan chan Runnable, closeHandle chan bool) *Worker {
	return &Worker{workerPool: workerPool, jobChannel: make(chan Runnable), closeHandle: closeHandle}
}

// Start starts the worker by listening to the job channel
func (w Worker) Start() {
	go func() {
		for {
			// Put the worker to the worker threadpool
			w.workerPool <- w.jobChannel

			select {
			// Wait for the job
			case job := <-w.jobChannel:
				// Got the job
				job.Run() //run the job
			case <-w.closeHandle:
				// Exit the go routine when the closeHandle channel is closed
				return
			}
		}
	}()
}
