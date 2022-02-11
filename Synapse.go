package main

import (
	"context"
	"fmt"
	"log"
	"synapse/advert"
	"synapse/constants"
	gui2 "synapse/gui"
	"synapse/logging"
)

func Init() {

	ctx, cn := context.WithCancel(context.Background())
	defer cn()

	logger, lgO, err := logging.New()
	//_, amO, err := amazonad.New()
	_, ht0, err := advert.NewHttp(ctx, logger)

	build := map[gui2.Identifier]*gui2.GuiIdentifier{
		gui2.LoggerId: lgO,
		"http":        ht0,
	}

	if err != nil {
		log.Fatalln(err)
	}

	logger.Info(fmt.Sprintf("%s loaded and ready to whoop some streamer ass!", constants.Ver))
	logger.Info("https://discord.com/superyuuki")
	logger.Error("Hello bitbuckets ;)")

	if err := gui2.InitGui(ctx, cn, build); err != nil {
		log.Fatalln(err)
	}
}
