package task

import "errors"

type ThreadPool struct {
	queueSize   int64
	noOfWorkers int

	jobQueue    chan Runnable
	workerPool  chan chan Runnable
	closeHandle chan bool // Channel used to stop all the workers
}

// NewThreadPool creates thread threadpool
func NewThreadPool(noOfWorkers int, queueSize int64) *ThreadPool {
	threadPool := &ThreadPool{queueSize: queueSize, noOfWorkers: noOfWorkers}
	threadPool.jobQueue = make(chan Runnable, queueSize)
	threadPool.workerPool = make(chan chan Runnable, noOfWorkers)
	threadPool.closeHandle = make(chan bool)
	threadPool.createPool()
	return threadPool
}

func NewPopulatedThreadPool(noOfWorkers int, queueSize int64, runnable Runnable) {

}

func (t *ThreadPool) Submit(task Runnable) error {
	if len(t.jobQueue) == int(t.queueSize) {
		return errors.New("ITS FULL BITCH")
	}
	t.jobQueue <- task
	return nil
}

// Close will close the threadpool
// It sends the stop signal to all the worker that are running
//TODO: need to check the existing /running task before closing the threadpool
func (t *ThreadPool) Close() {
	close(t.closeHandle) // Stops all the routines
	close(t.workerPool)  // Closes the Job threadpool
	close(t.jobQueue)    // Closes the job Queue
}

// createPool creates the workers and start listening on the jobQueue
func (t *ThreadPool) createPool() {
	for i := 0; i < t.noOfWorkers; i++ {
		worker := NewWorker(t.workerPool, t.closeHandle)
		worker.Start()
	}

	go t.dispatch()

}

// dispatch listens to the jobqueue and handles the jobs to the workers
func (t *ThreadPool) dispatch() {
	for {
		select {

		case job := <-t.jobQueue:
			// Got job
			func(job Runnable) {
				//Find a worker for the job
				jobChannel := <-t.workerPool
				//Submit job to the worker
				jobChannel <- job
			}(job)

		case <-t.closeHandle:
			// Close thread threadpool
			return
		}
	}
}
