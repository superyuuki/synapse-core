package gui

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/button"
)

const LoggerId = "logger"

type GuiImpl struct {
	innerContainer *container.Container
	buttonBar      []container.Option

	containerMap map[Identifier][]container.Option
}

func (m GuiImpl) Mux(identifier Identifier) error {

	if identifier == LoggerId {
		return m.innerContainer.Update(containerId,
			container.SplitHorizontal(
				container.Top(m.buttonBar...),
				container.Bottom(m.containerMap[LoggerId]...),
				container.SplitPercent(10),
			),
		)
	}

	return m.innerContainer.Update(containerId,
		container.SplitHorizontal(
			container.Top(m.buttonBar...),
			container.Bottom(
				container.SplitVertical(
					container.Left(m.containerMap[LoggerId]...),
					container.Right(m.containerMap[identifier]...),
					container.SplitPercent(35),
				),
			),
			container.SplitPercent(10),
		),
	)

}

func NewGui(cnt *container.Container, inmap map[Identifier]*GuiIdentifier) (Gui, error) {

	localMap := make(map[Identifier][]container.Option)

	preGui := &GuiImpl{
		innerContainer: cnt,
		buttonBar:      nil,
		containerMap:   localMap,
	}

	percent := 99 / len(inmap)
	builder := grid.New()

	for id, data := range inmap {
		str := id

		localMap[str] = data.opts

		btn, err := button.New(data.displayName, func() error {
			return preGui.Mux(str)
		},
			button.GlobalKey(data.key),
			button.Height(1),
			button.Width(30),
			button.DisableShadow(),
			button.FillColor(cell.ColorYellow),
			button.PressedFillColor(cell.ColorRed),
		)

		if err != nil {
			return nil, err
		}

		builder.Add(
			grid.ColWidthPerc(
				percent, grid.Widget(btn, container.PaddingRightPercent(0), container.PaddingLeftPercent(0)),
			),
		)

	}

	top, err := builder.Build()

	if err != nil {
		return nil, err
	}

	top = append(top,
		container.PaddingBottomPercent(0),
		container.PaddingLeftPercent(0),
		container.PaddingRightPercent(0),
		container.PaddingTopPercent(0),
		container.Border(linestyle.Round),
		container.BorderTitle("Menu"),
	)

	preGui.buttonBar = top

	return preGui, nil

}
