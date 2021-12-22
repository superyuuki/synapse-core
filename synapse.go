package main

import (
	"fmt"
	"log"
	"synapse/advert"
	"synapse/constants"
	"synapse/gui"
	"synapse/logging"
)

func main() {

	logger, lgO, err := logging.New()
	_, amO, err := advert.New()
	seO, err := constants.New()
	hpO, err := constants.New()

	if err != nil {
		log.Fatalln(err)
	}

	//init starting values, etc etc

	logger.Info(fmt.Sprintf("%s loaded and ready to whoop some streamer ass!", constants.Ver))
	logger.Info("https://discord.com/superyuuki")
	logger.Error("Hello bitbuckets ;)")

	if err := gui.InitGui(lgO, amO, seO, hpO); err != nil {
		log.Fatalln(err)
	}

}
