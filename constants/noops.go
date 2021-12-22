package constants

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/text"
)

func New() ([]container.Option, error) {
	txt, err := text.New()

	if err != nil {
		return nil, err
	}

	if err := txt.Write("Non Implemented Panel"); err != nil {
		return nil, err
	}

	cnt := []container.Option{container.PlaceWidget(txt), container.Border(linestyle.Light), container.BorderTitle("Debug")}

	return cnt, nil
}
