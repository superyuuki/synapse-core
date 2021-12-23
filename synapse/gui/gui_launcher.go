package gui

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"synapse/synapse/constants"
)

const containerId = "rootId"

func InitGui(ctx context.Context, cn func(), vmap map[Identifier]*GuiIdentifier) error {

	t, err := tcell.New()
	if err != nil {
		return err
	}
	defer t.Close()

	c, err := container.New(
		t,
		container.ID(containerId),
		container.Border(linestyle.Light),
		container.BorderTitle(fmt.Sprintf("%s | Press Q to quit", constants.Ver)),
	)

	if err != nil {
		return err
	}

	headboard, err := NewGui(c, vmap)

	if err != nil {
		return err
	}

	if err := headboard.Mux(LoggerId); err != nil {
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
