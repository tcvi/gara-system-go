package vehicleorder

import (
	"garasystem/internal/core/ports"
	"garasystem/internal/core/services"
)

type Service struct {
	repo        *services.Repository
	userService ports.UserService
}

func NewVehicleService(repo *services.Repository, userService ports.UserService) *Service {
	return &Service{repo: repo, userService: userService}
}
