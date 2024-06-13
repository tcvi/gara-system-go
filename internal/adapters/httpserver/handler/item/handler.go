package item

import "garasystem/internal/core/ports"

type Handler struct {
	service ports.ItemService
}

func NewHandler(service ports.ItemService) *Handler {
	return &Handler{service}
}
