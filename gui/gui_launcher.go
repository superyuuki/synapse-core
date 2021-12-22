package gui

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"synapse/constants"
)

func InitGui(lg []container.Option, am []container.Option, se []container.Option, hp []container.Option) error {

	t, err := tcell.New()
	if err != nil {
		return err
	}
	defer t.Close()

	ctx, cn := context.WithCancel(context.Background())
	defer cn()

	c, err := container.New(
		t,
		container.ID(containerId),
		container.Border(linestyle.Light),
		container.BorderTitle(fmt.Sprintf("%s | Press q to quit", constants.Ver)),
	)

	if err != nil {
		return err
	}

	headboard, err := newGui(c, lg, am, se, hp)

	if err != nil {
		return err
	}

	if err := headboard.SetLogger(); err != nil {
		return err
	}

	quitter := func(keyboard *terminalapi.Keyboard) {
		input := keyboard.Key.String()

		if input == "q" || input == "Q" || input == "KeyEsc" || input == "KeyCtrlC" {
			cn()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
		return err
	}

	return nil

}
