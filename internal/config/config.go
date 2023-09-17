package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type configKey string

const (
	GrpcPORT = configKey("GRPC_PORT")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("error while reading env file")
	}
}

func GetValue(key configKey) interface{} {
	return os.Getenv(string(key))
}
