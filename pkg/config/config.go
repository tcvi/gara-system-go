package config

import (
	"context"
	"fmt"
	"garasystem/internal/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/lpernett/godotenv"
)

var config Config

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

	Hook struct {
		Mattermost string `envconfig:"MATTERMOST_URL"`
	}

	AWS struct {
		BucketName string `envconfig:"AWS_BUCKET_NAME"`
	}
}

func LoadConfig() (*Config, error) {
	// load default .env file, ignore the error
	_ = godotenv.Load()

	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("load config error: %v", err)
	}

	return &config, nil
}

func GetConfig() *Config {
	return &config
}

func LoadAwsConfig(cfg *Config) aws.Config {
	awsConfig, err := awsconfig.LoadDefaultConfig(
		context.TODO(),
		awsconfig.WithRegion("ap-southeast-1"),
		awsconfig.WithSharedConfigProfile("gara-system-dev"),
	)
	if err != nil {
		logger.Log.Fatal("fail to load aws config: ", err)
	}

	return awsConfig
}
