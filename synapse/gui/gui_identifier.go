package gui

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
)

type GuiIdentifier struct {
	displayName string
	key         keyboard.Key
	opts        []container.Option
}

func NewIdentifier(displayName string, key keyboard.Key, opts []container.Option) *GuiIdentifier {
	return &GuiIdentifier{
		displayName: displayName,
		key:         key,
		opts:        opts,
	}
}
