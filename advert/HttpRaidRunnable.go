package advert

import (
	"synapse/logging"
	"time"
)

type HttpRaidRunnable struct {
	lg logging.Logger
}

func NewHttpRaidRunnable(lg logging.Logger) *HttpRaidRunnable {
	return &HttpRaidRunnable{lg: lg}
}

func (h *HttpRaidRunnable) Run() {
	time.Sleep(5 * time.Second)

	h.lg.Info("Hi")
}
