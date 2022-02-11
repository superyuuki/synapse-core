package newtype

import "context"

type RunImpl struct {
	holderContext context.Context
	holderCancel  func()
}

func (t *RunImpl) Run() {
	go func() {

		for {

		}

	}()
}

func (t *RunImpl) doSomething(completion chan int) {

	completion <- 1
}
