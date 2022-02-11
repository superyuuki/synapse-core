package advert

import (
	"github.com/shettyh/threadpool"
	"synapse/task"
	"sync"
)

type TaskerRunnable struct {
	wg *sync.WaitGroup

	delegate threadpool.Runnable
}

func NewTaskerRunnable(wg *sync.WaitGroup, runnable task.Runnable) *TaskerRunnable {
	return &TaskerRunnable{wg: wg, delegate: runnable}
}

func (t *TaskerRunnable) Run() {
	defer t.wg.Done()

	t.delegate.Run()
}
