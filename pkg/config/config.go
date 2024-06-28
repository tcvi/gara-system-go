package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/lpernett/godotenv"
)

type Config struct {
	AppEnv       string `envconfig:"APP_ENV"`
	Port         int    `envconfig:"PORT"`
	JwtSecretKey string `envconfig:"JWT_SECRET_KEY"`

	DB struct {
		Name      string `envconfig:"DB_NAME"`
		Host      string `envconfig:"DB_HOST"`
		Port      int    `envconfig:"DB_PORT"`
		User      string `envconfig:"DB_USER"`
		Pass      string `envconfig:"DB_PASS"`
		EnableSSL bool   `envconfig:"ENABLE_SSL"`
	}

	Redis struct {
		Password string `envconfig:"REDIS_PASSWORD"`
		Port     string `envconfig:"REDIS_PORT"`
		Host     string `envconfig:"REDIS_HOST"`
	}
}

func LoadConfig() (*Config, error) {
	// load default .env file, ignore the error
	_ = godotenv.Load()

	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("load config error: %v", err)
	}

	return cfg, nil
}
