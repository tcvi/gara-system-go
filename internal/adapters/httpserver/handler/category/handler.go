package category

import "garasystem/internal/core/ports"

type Handler struct {
	service ports.CategoryService
}

func NewHandler(service ports.CategoryService) *Handler {
	return &Handler{service}
}
