package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type configKey string

const (
	GrpcPORT         = configKey("GRPC_PORT")
	JWTTokenDuration = configKey("JWT_TOKEN_DURATION")
)

type PostgresConfig struct {
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDb       string `env:"POSTGRES_DB"`
	PostgresDbHost   string `env:"POSTGRES_DB_HOST"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
}

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

func GetPostgresConfig() *PostgresConfig {
	cfg := &PostgresConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Warn().Err(err).Msgf("unable to parse PostgresConfig")
	}

	return cfg
}

func GetJWTTokenDuration() (time.Duration, error) {
	tokenDuration := os.Getenv(string(JWTTokenDuration))
	tokenDurationTime, err := time.ParseDuration(tokenDuration)
	if err != nil {
		return time.Duration(0), err
	}

	return tokenDurationTime, nil
}
