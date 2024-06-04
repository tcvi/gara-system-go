package postgrestorage

import (
	"fmt"
	"strconv"

	"garasystem/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Options struct {
	DBName   string
	DBUser   string
	Password string
	Host     string
	Port     string
	SSLMode  bool
}

func ParseFromConfig(c *config.Config) Options {
	return Options{
		DBName:   c.DB.Name,
		DBUser:   c.DB.User,
		Password: c.DB.Pass,
		Host:     c.DB.Host,
		Port:     strconv.Itoa(c.DB.Port),
		SSLMode:  c.DB.EnableSSL,
	}
}

func NewConnection(opts Options) (db *gorm.DB, err error) {
	sslmode := "disable"
	if opts.SSLMode {
		sslmode = "enable"
	}

	datasource := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		opts.Host, opts.Port, opts.DBUser, opts.Password, opts.DBName, sslmode,
	)
	return gorm.Open(postgres.Open(datasource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
