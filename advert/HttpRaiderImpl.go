package advert

import (
	"errors"
	"github.com/shettyh/threadpool"
	"synapse/logging"
	"sync"
)

type HttpRaiderImpl struct {
	mx *sync.Mutex
	lg logging.Logger

	//run specific
	tp *threadpool.ThreadPool //nullable
	wg *sync.WaitGroup
}

func (t *HttpRaiderImpl) StartRaid(maxConcurrent int, maxTasks int) error {

	t.mx.Lock()
	defer t.mx.Unlock()

	if t.tp != nil {
		return errors.New("the HttpRaider is already running")
	}

	t.tp = threadpool.NewThreadPool(maxConcurrent, int64(maxTasks))
	t.wg = &sync.WaitGroup{} //new waitgroup

	newRunnable := NewHttpRaidRunnable(t.lg)

	//fill that shit with misery
	for i := 0; i < maxTasks; i++ {
		err := t.tp.Execute(NewTaskerRunnable(t.wg, newRunnable))
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *HttpRaiderImpl) FinishRaid() error {
	t.mx.Lock()
	defer t.mx.Unlock()

	if tp == nil {
		return errors.New("the HttpRaider is not running")
	}

	tp.Close()

	return nil
}

func (t *HttpRaiderImpl) CancelRaid() error {

}
