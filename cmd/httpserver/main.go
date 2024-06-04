package main

import (
	"fmt"
	"garasystem/internal/adapters/aws"
	"garasystem/internal/adapters/aws/snsservice"
	"garasystem/internal/adapters/httpserver"
	"garasystem/internal/adapters/postgrestorage"
	userstorage "garasystem/internal/adapters/postgrestorage/user"
	"garasystem/internal/core/services"
	userservice "garasystem/internal/core/services/user"
	"garasystem/internal/logger"
	"garasystem/pkg/config"
	"net/http"
)

var (
	userService *userservice.Service
)

func init() {
	logger.SetupLogger()
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.Fatal(err)
	}

	awsConfig, err := aws.LoadConfig()
	if err != nil {
		logger.Log.Fatal(err)
	}

	db, err := postgrestorage.NewConnection(postgrestorage.ParseFromConfig(cfg))
	//db, err := dynamodbstorage.NewConnection(awsConfig)
	if err != nil {
		logger.Log.Fatal(err)
	}

	userStore := userstorage.NewStorage(db)
	//userStore := userstorage.NewStorage(db)

	repo := services.NewRepository(userStore)

	snsService := snsservice.NewSnsService(awsConfig)
	userService = userservice.NewUserService(repo, snsService)

	server, _ := httpserver.NewServer(userService, snsService)

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Log.Println("server started at port", addr)
	logger.Log.Fatal(http.ListenAndServe(addr, server))
}
