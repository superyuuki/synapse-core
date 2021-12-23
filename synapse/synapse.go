package synapse

import (
	"context"
	"fmt"
	"log"
	"synapse/synapse/advert/httpad"
	"synapse/synapse/constants"
	"synapse/synapse/gui"
	"synapse/synapse/logging"
)

func Init() {

	ctx, cn := context.WithCancel(context.Background())
	defer cn()

	logger, lgO, err := logging.New()
	//_, amO, err := amazonad.New()
	_, ht0, err := httpad.NewHttp(ctx, logger)

	build := map[gui.Identifier]*gui.GuiIdentifier{
		gui.LoggerId: lgO,
		"httpad":     ht0,
	}

	if err != nil {
		log.Fatalln(err)
	}

	logger.Info(fmt.Sprintf("%s loaded and ready to whoop some streamer ass!", constants.Ver))
	logger.Info("https://discord.com/superyuuki")
	logger.Error("Hello bitbuckets ;)")

	if err := gui.InitGui(ctx, cn, build); err != nil {
		log.Fatalln(err)
	}
}
