package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/do/v2"
)

type Config struct {
	Version string `env:"VERSION" env-default:"-"`
	Cors    string `env:"CORS"`
}

func NewConfig(i do.Injector) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func init() {
	do.Provide(nil, NewConfig)
}
