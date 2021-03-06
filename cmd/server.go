package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/alexferl/go_grpc_helloworld"
)

func main() {
	c := go_grpc_helloworld.NewConfig()
	c.BindFlags()

	log.Info().Msgf("Starting environment: %s", viper.GetString("env-name"))
	go_grpc_helloworld.Start()
}
