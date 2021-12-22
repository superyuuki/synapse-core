package advert

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
)

type AdvertLambda struct {
	chart *linechart.LineChart
}

func (a AdvertLambda) StartRaid() {

	//TODO implement me
	panic("implement me")
}

func (a AdvertLambda) StopRaid() {
	//TODO implement me
	panic("implement me")
}

func New() (Advert, []container.Option, error) {
	chart, err := linechart.New()

	if err != nil {
		return nil, nil, err
	}

	tx, err := text.New()

	if err != nil {
		return nil, nil, err
	}

	if err := tx.Write("Please put the targetting form here, thanks"); err != nil {
		return nil, nil, err
	}

	advert := &AdvertLambda{chart: chart}

	strc := container.SplitVertical(
		container.Left(
			container.PlaceWidget(chart),
		),
		container.Right(
			container.PlaceWidget(tx),
		),
	)

	cnt := []container.Option{
		strc,
		container.BorderTitle("Amazon Lambda Fucker"),
		container.Border(linestyle.Light),
	}

	return advert, cnt, nil
}
