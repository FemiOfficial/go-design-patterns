package config

import (
	"github.com/caarlos0/env/v10"
)

type Configuration struct {
	OpenWeatherAPIKey string `env:"OPENWEATHER_API_KEY"`
}

func NewConfig() (*Configuration, error) {
	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}