package gui

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
)

type GuiImpl struct {
	container *container.Container

	logger   []container.Option
	amazon   []container.Option
	selenium []container.Option
	help     []container.Option

	buttons *buttons
}

func (h GuiImpl) set(btm []container.Option) error {

	return h.container.Update(containerId,
		container.SplitHorizontal(
			container.Top(
				container.SplitVertical(
					container.Left(
						container.SplitVertical(
							container.Left(container.PlaceWidget(h.buttons.loggerButton)),
							container.Right(container.PlaceWidget(h.buttons.helpButton)),
						),
					),
					container.Right(
						container.SplitVertical(
							container.Left(container.PlaceWidget(h.buttons.amazonButton)),
							container.Right(container.PlaceWidget(h.buttons.seleniumButton)),
						),
					),
				),
				container.Border(linestyle.Round),
				container.BorderTitle("Menu"),
			),
			container.Bottom(btm...),
			container.SplitPercent(20),
		),
	)
}

func (h GuiImpl) SetLogger() error {
	return h.set(h.logger)
}

func (h GuiImpl) SetAmazonBotter() error {
	return h.set([]container.Option{
		container.SplitVertical(
			container.Left(h.logger...),
			container.Right(h.amazon...),
		),
	})
}

func (h GuiImpl) SetSeleniumBotter() error {
	return h.set([]container.Option{
		container.SplitVertical(
			container.Left(h.logger...),
			container.Right(h.selenium...),
		),
	})
}
func (h GuiImpl) SetHelp() error {
	return h.set([]container.Option{
		container.SplitVertical(
			container.Left(h.logger...),
			container.Right(h.help...),
		),
	})
}

func newGui(container2 *container.Container, lg []container.Option, am []container.Option, se []container.Option, hp []container.Option) (*GuiImpl, error) {
	preboard := &GuiImpl{
		container: container2,
		logger:    lg,
		amazon:    am,
		selenium:  se,
		help:      hp,
	}

	buttons, err := newButtons(preboard)

	if err != nil {
		return nil, err
	}

	preboard.buttons = buttons

	return preboard, nil
}
