package main

import (
	"fmt"
	"garasystem/internal/adapters/aws"
	"garasystem/internal/adapters/aws/snsservice"
	"garasystem/internal/adapters/httpserver"
	"garasystem/internal/adapters/postgrestorage"
	categorystorage "garasystem/internal/adapters/postgrestorage/category"
	userstorage "garasystem/internal/adapters/postgrestorage/user"
	vehicleorderstorage "garasystem/internal/adapters/postgrestorage/vehicleorder"
	"garasystem/internal/core/services"
	categoryservice "garasystem/internal/core/services/category"
	userservice "garasystem/internal/core/services/user"
	vehicleorderservice "garasystem/internal/core/services/vehicleorder"
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
	vehicleStore := vehicleorderstorage.NewStorage(db)
	categoryStore := categorystorage.NewStorage(db)

	repo := services.NewRepository(userStore, vehicleStore, categoryStore)

	snsService := snsservice.NewSnsService(awsConfig)
	userService = userservice.NewUserService(repo, snsService)
	vehicleOrderService := vehicleorderservice.NewVehicleService(repo, userService)
	categoryService := categoryservice.NewService(repo)

	server, _ := httpserver.NewServer(cfg, userService, vehicleOrderService, categoryService, snsService)

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Log.Println("server started at port", addr)
	logger.Log.Fatal(http.ListenAndServe(addr, server))
}
