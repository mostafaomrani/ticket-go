package main

import (
	"ticket/internal/utils"

	"github.com/rs/zerolog/log"
)

func main() {

	// Load config
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Can`t load config file")
	}

	if config.APPDEBUG == "true" {
		log.Info().Str("foo", "bar").Msg("Hello world")
	}

}
