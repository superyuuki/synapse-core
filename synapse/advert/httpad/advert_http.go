package httpad

import (
	"context"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/linechart"
	"strconv"
	"synapse/synapse/advert"
	"synapse/synapse/advert/pool"
	"synapse/synapse/gui"
	"synapse/synapse/logging"
	"sync"
)

type AdvertHttp struct {
	logger logging.Logger //
	chart  *linechart.LineChart

	start chan *pool.StartSignal
	stop  chan pool.Termination
}

func NewHttp(ctx context.Context, logger logging.Logger) (advert.Advert, *gui.GuiIdentifier, error) {

	start := make(chan *pool.StartSignal)
	stop := make(chan pool.Termination)

	chart, err := linechart.New()
	perc, err := barchart.New()
	form, err := newForm(start, stop)

	if err != nil {
		return nil, nil, err
	}

	mxi := gui.NewIdentifier("(h)ttp", 'h', []container.Option{
		container.Border(linestyle.Double),
		container.BorderTitle("[HTTP] Twitch Ad Fucker"),
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(container.PlaceWidget(chart), container.Border(linestyle.Light)),
					container.Right(container.PlaceWidget(perc), container.Border(linestyle.Light)),
				),
			),
			container.Bottom(
				form...,
			),
		),
	})

	go InitPoolReceiver(ctx, logger, start, stop)

	return &AdvertHttp{
		logger: logger,
		chart:  chart,
		start:  start,
		stop:   stop,
	}, mxi, nil
}

func InitPoolReceiver(ctx context.Context, lg logging.Logger, start chan *pool.StartSignal, stop chan pool.Termination) {
	isRunning := false

	for {
		select {
		case <-ctx.Done():
			stop <- pool.FailedTerminated //forward
		case sig := <-stop:
			if isRunning {

				isRunning = false

				switch sig {
				case pool.FailedTerminated:
					lg.Error("Stopped: user stopped signal")

				case pool.FailedExceptional:
					lg.Error("Stopped: something fucked")

				case pool.Success:
					lg.Info("Stopped: successful raid enjoy twitch prime lol")
				}

			} else {
				lg.Error("Cannot stop: already stopped!")
			}

		case sig := <-start:
			if !isRunning {
				isRunning = true

				go InitRaid(sig, lg, stop)
			} else {
				lg.Error("Cannot start: already running!")
			}
		}
	}

}

func InitRaid(sig *pool.StartSignal, lg logging.Logger, stop chan pool.Termination) {

	lg.Info("Starting raid")

	input := make(chan pool.InSignal, sig.MaxTasks)
	output := make(chan pool.OutSignal, sig.MaxTasks)
	wg := &sync.WaitGroup{}

	delta := 0

	for i := 0; i < sig.MaxWorkers; i++ {
		wg.Add(1)

		delta = delta + 1
		go pool.Worker(lg, input, output, wg)
	}

	//populate input channel
	for i := 0; i < sig.MaxTasks; i++ {
		input <- pool.InSignal{
			Target: sig.Target + " iteration " + strconv.Itoa(i),
		}
	}

	go func() {
		// Close input channel since no more jobs are being sent to input channel
		close(input)

		// Wait for all goroutines to finish processing
		wg.Wait() //should be reduced if termination signal is sent anyways

		// Close output channel since all workers have finished processing
		close(output)

	}()

	stop <- pool.Success
}

func (a AdvertHttp) StartRaid() error {

	go func() {
		a.start <- &pool.StartSignal{}
	}()

	return nil

}

func (a AdvertHttp) StopRaid() error {
	go func() {
		a.stop <- pool.FailedTerminated
	}()

	return nil
}
