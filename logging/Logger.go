package logging

import "github.com/mum4k/termdash/cell"

type Logger interface {
	Error(details string)
	Info(details string)
	Prefixed(prefix string, details string, color cell.Color)
}
