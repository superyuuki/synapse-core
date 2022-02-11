package advert

import (
	"context"
	"github.com/shettyh/threadpool"
	"sync"
)

type CancellablePool struct {
	ctx    context.Context //parent
	cancel context.CancelFunc

	maxConcurrent int
	maxTasks      int
}

func (t *CancellablePool) Start() {

	//pop that waitgroup up!!!!

	tp := threadpool.NewThreadPool(t.maxConcurrent, int64(t.maxTasks))
	wg := &sync.WaitGroup{} //new waitgroup

	newRunnable := NewHttpRaidRunnable(nil)

	//fill that shit with misery
	for i := 0; i < t.maxTasks; i++ {
		err := tp.Execute(NewTaskerRunnable(wg, newRunnable))
		if err != nil {
			panic(err)
		}
	}

	for {
		select {
		case <-t.ctx.Done():

		}
	}
