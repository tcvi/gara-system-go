package vehicleorder

import (
	"garasystem/internal/core/ports"
	"garasystem/internal/core/services"
)

type Service struct {
	repo                    *services.Repository
	userService             ports.UserService
	vehicleOrderItemService ports.VehicleOrderItemService
}

func NewVehicleService(
	repo *services.Repository,
	userService ports.UserService,
	vehicleOrderItemService ports.VehicleOrderItemService,
) *Service {
	return &Service{
		repo:                    repo,
		userService:             userService,
		vehicleOrderItemService: vehicleOrderItemService,
	}
}
