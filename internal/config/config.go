package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AppPort     int
	DatabaseURI string
}

func (c *Config) ValidateEnv() error {
	if c.AppPort < 1 || c.AppPort > 65536 {
		return errors.New("invalid APP_PORT: must be between 1 and 65535")
	}

	if c.DatabaseURI == "" {
		return errors.New("MONGO_URI is a required configuration field")
	}

	return nil
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Load env")
	}

	if appEnv := os.Getenv("APP_ENV"); appEnv == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
		})
	}

	portStr := os.Getenv("APP_PORT")
	if portStr == "" {
		return nil, errors.New("APP_PORT environment variable is missing")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("APP_PORT must be a valid integer")
	}

	cfg := &Config{
		AppPort:     port,
		DatabaseURI: os.Getenv("MONGO_URI"),
	}

	if err = cfg.ValidateEnv(); err != nil {
		return nil, err
	}

	return cfg, nil
}
