package gui

const containerId = "rootId"

type Gui interface {
	SetLogger() error
	SetAmazonBotter() error
	SetSeleniumBotter() error
	SetHelp() error
}
