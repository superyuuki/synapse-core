package logging

import (
	"fmt"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/text"
	"synapse/synapse/gui"
	"time"
)

type LoggerImpl struct {
	text *text.Text
}

func (l LoggerImpl) Error(details string) {

	err := l.text.Write(fmt.Sprintf("[%s] [Error] %s \n", time.Now().Format("2006-01-02 15:04:05"), details), text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
	if err != nil {
		panic(err)
	}
}

func (l LoggerImpl) Info(details string) {
	err := l.text.Write(fmt.Sprintf("[%s] [Info] %s \n", time.Now().Format("2006-01-02 15:04:05"), details), text.WriteCellOpts(cell.FgColor(cell.ColorSilver)))
	if err != nil {
		panic(err)
	}
}

func (l LoggerImpl) Prefixed(prefix string, details string, color cell.Color) {
	err := l.text.Write(fmt.Sprintf("[%s] [%s] %s \n", time.Now().Format("2006-01-02 15:04:05"), prefix, details), text.WriteCellOpts(cell.FgColor(color)))
	if err != nil {
		panic(err)
	}
}

func New() (Logger, *gui.GuiIdentifier, error) {

	txt, err := text.New(text.RollContent(), text.WrapAtWords())

	if err != nil {
		return nil, nil, err
	}

	logger := &LoggerImpl{text: txt}

	mxi := gui.NewIdentifier(
		"(l)ogger",
		'l',
		[]container.Option{
			container.PlaceWidget(txt),
			container.Border(linestyle.Light),
			container.BorderTitle("Logs"),
		},
	)

	return logger, mxi, nil

}
