package gui

type Identifier string

type Gui interface {
	Mux(identifier Identifier) error
}
