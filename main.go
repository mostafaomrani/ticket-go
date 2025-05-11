package main

import (
	"context"
	"fmt"
	"ticket/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
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

	// Connect to the database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUSERNAME,
		config.DBPASSWORD,
		config.DBHOST,
		config.DBPORT,
		config.DBDATABASE,
	)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create connection pool")
	}
	defer dbpool.Close()

}
