package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"log"
)

type twitchConfig struct {
	twitchAddress string `mapstructure:"twitchAddress"`
}

type driver interface {
}

func main() {
	log.Println("[INIT] Launching Synapse by Yuuki")

	log.Println("[INIT] Loading configuration")

	config.AddDriver(yaml.Driver)
	twitchConfig := twitchConfig{}

	err := config.BindStruct("../config.yml", &twitchConfig)

	if err != nil {
		panic(err)
	}

	log.Println(twitchConfig.twitchAddress)
}
