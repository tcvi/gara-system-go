package main

import (
	"fmt"
	"garasystem/internal/adapters/aws"
	"garasystem/internal/adapters/aws/snsservice"
	"garasystem/internal/adapters/httpserver"
	"garasystem/internal/adapters/postgrestorage"
	categorystorage "garasystem/internal/adapters/postgrestorage/category"
	itemstorage "garasystem/internal/adapters/postgrestorage/item"
	userstorage "garasystem/internal/adapters/postgrestorage/user"
	vehicleorderstorage "garasystem/internal/adapters/postgrestorage/vehicleorder"
	vehicleorderitemstorage "garasystem/internal/adapters/postgrestorage/vehicleorderitem"
	"garasystem/internal/adapters/redis"
	"garasystem/internal/core/services"
	categoryservice "garasystem/internal/core/services/category"
	itemservice "garasystem/internal/core/services/item"
	notificationservice "garasystem/internal/core/services/notification"
	"garasystem/internal/core/services/redistask"
	userservice "garasystem/internal/core/services/user"
	vehicleorderservice "garasystem/internal/core/services/vehicleorder"
	vehicleorderitemservice "garasystem/internal/core/services/vehicleorderitem"
	"garasystem/internal/logger"
	"garasystem/pkg/config"
	"net/http"
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
	if err != nil {
		logger.Log.Fatal(err)
	}

	userStore := userstorage.NewStorage(db)
	vehicleStore := vehicleorderstorage.NewStorage(db)
	categoryStore := categorystorage.NewStorage(db)
	itemStore := itemstorage.NewStorage(db)
	vehicleOrderItemsStore := vehicleorderitemstorage.NewStorage(db)

	repo := services.NewRepository(userStore, vehicleStore, categoryStore, itemStore, vehicleOrderItemsStore)

	snsService := snsservice.NewSnsService(awsConfig)
	userService := userservice.NewUserService(repo, snsService)
	itemService := itemservice.NewService(repo)
	vehicleOrderItemService := vehicleorderitemservice.NewVehicleService(repo, itemService)
	vehicleOrderService := vehicleorderservice.NewVehicleService(repo, userService, vehicleOrderItemService)
	categoryService := categoryservice.NewService(repo)
	notificationService, err := notificationservice.NewService()
	if err != nil {
		logger.Log.Fatal("Create notificationService fail ", err)
	}
	redisClient := redistask.NewRedisTaskClient(cfg)
	server := httpserver.NewServer(cfg,
		userService,
		vehicleOrderService,
		categoryService,
		itemService,
		vehicleOrderItemService,
		notificationService,
		redisClient,
		snsService,
	)

	// Start redis task server
	go func() {
		redis.NewServer(cfg, notificationService)
	}()

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Log.Println("server started at port", addr)
	logger.Log.Fatal(http.ListenAndServe(addr, server))
}
