package vehicleoderitem

import "garasystem/internal/core/ports"

type Handler struct {
	service ports.VehicleOrderService
}

func NewHandler(service ports.VehicleOrderService) *Handler {
	return &Handler{service: service}
}
