package advert

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/linechart"
	"synapse/gui"
	"synapse/logging"
	"time"
)

type HttpApi struct {
}

func (a HttpApi) StartRaid() error {
	//TODO
	panic("impl")
}

func (a HttpApi) StopRaid() error {
	//TODO
	panic("impl")
}

func NewHttp(ctx context.Context, lg logging.Logger) (HttpRaider, *gui.GuiIdentifier, error) {

	processor := func(i pool.Input) {
		lg.Prefixed("start", fmt.Sprintf("Starting: %s", i), cell.ColorGreen)
		time.Sleep(time.Second)
		lg.Prefixed("stop", fmt.Sprintf("Stopping: %s", i), cell.ColorRed)
	}

	ch := pool.NewBase(processor, lg).Init()

	chart, err := linechart.New()
	perc, err := barchart.New()
	form, err := newForm(ch)

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

	return HttpApi{}, mxi, nil

}
