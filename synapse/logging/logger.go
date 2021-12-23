package logging

import "github.com/mum4k/termdash/cell"

const Cum = ""

type Logger interface {
	Error(details string)
	Info(details string)
	Prefixed(prefix string, details string, color cell.Color)
}
