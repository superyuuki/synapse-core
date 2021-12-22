package gui

import (
	"github.com/mum4k/termdash/widgets/button"
)

type buttons struct {
	loggerButton   *button.Button
	amazonButton   *button.Button
	seleniumButton *button.Button
	helpButton     *button.Button
}

func newButtons(board *GuiImpl) (*buttons, error) {
	loggerButton, err := button.New("(l)ogger", func() error {
		return board.SetLogger()
	},

		button.GlobalKey('l'),
	)

	amazonButton, err := button.New("(a)mazon", func() error {
		return board.SetAmazonBotter()
	},

		button.GlobalKey('a'),
	)

	selenium, err := button.New("(s)elenium", func() error {
		return board.SetSeleniumBotter()
	},

		button.GlobalKey('s'),
	)

	help, err := button.New("(h)elp", func() error {
		return board.SetHelp()
	},

		button.GlobalKey('h'),
	)

	if err != nil {
		return nil, err
	}

	return &buttons{
		loggerButton:   loggerButton,
		amazonButton:   amazonButton,
		seleniumButton: selenium,
		helpButton:     help,
	}, nil
}
