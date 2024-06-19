package vehicleorderitem

import (
	"garasystem/internal/core/ports"
	"garasystem/internal/core/services"
)

type Service struct {
	repo        *services.Repository
	itemService ports.ItemService
}

func NewVehicleService(repo *services.Repository, itemService ports.ItemService) *Service {
	return &Service{repo: repo, itemService: itemService}
}
