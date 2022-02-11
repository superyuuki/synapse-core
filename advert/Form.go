package advert

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
)

func newForm(localStop chan<- pool.BaseSignal) ([]container.Option, error) {
	startButton, err := button.New("Start the Fucker", func() error {
		go func() {
			localStop <- pool.Start
		}()

		return nil
	})

	stopButton, err := button.New("Stop the Fucker", func() error {
		go func() {
			localStop <- pool.Stop
		}()

		return nil
	})

	if err != nil {
		return nil, err
	}

	return []container.Option{

		container.SplitHorizontal(
			container.Top(
				container.PlaceWidget(startButton),
			),
			container.Bottom(
				container.PlaceWidget(stopButton),
			),
		),

		container.Border(linestyle.Light),
	}, nil
}
