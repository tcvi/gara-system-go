package vehicleoderitem

import "garasystem/internal/core/ports"

type Handler struct {
	service ports.VehicleOrderItemService
}

func NewHandler(service ports.VehicleOrderItemService) *Handler {
	return &Handler{service: service}
}
