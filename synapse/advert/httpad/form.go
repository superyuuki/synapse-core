package httpad

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
	"synapse/synapse/advert/pool"
)

func newForm(start chan *pool.StartSignal, stop chan pool.Termination) ([]container.Option, error) {
	startbtn, err := button.New("start", func() error {
		go func() {
			start <- &pool.StartSignal{
				Target:     "cum",
				MaxWorkers: 10,
				MaxTasks:   100,
			}
		}()

		return nil
	})

	stopbtn, err := button.New("stop", func() error {
		go func() {
			stop <- pool.FailedTerminated
		}()

		return nil
	})

	if err != nil {
		return nil, err
	}

	return []container.Option{

		container.SplitHorizontal(
			container.Top(
				container.PlaceWidget(startbtn),
			),
			container.Bottom(
				container.PlaceWidget(stopbtn),
			),
		),

		container.Border(linestyle.Light),
	}, nil
}
