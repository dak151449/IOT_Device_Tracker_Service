package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

type configKey string

const (
	GrpcPORT = configKey("GRPC_PORT")
)

func init() {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("error while reading env file")
		}
	}
}

func VariableNotFoundMsg(key configKey) string {
	return fmt.Sprintf("secret %s not found", key)
}

func GetStringValue(key configKey) string {
	value, exists := os.LookupEnv(string(key))
	if !exists {
		log.Warn().Msg(VariableNotFoundMsg(key))
		return ""
	}

	return value
}

func GetIntValue(key configKey) int {
	value, exists := os.LookupEnv(string(key))
	if !exists {
		log.Warn().Msg(VariableNotFoundMsg(key))
		return 0
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert secret %s to int", key)
		return 0
	}

	return intValue
}
