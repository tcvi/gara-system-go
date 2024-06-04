package user

import (
	"garasystem/internal/core/ports"
)

type Handler struct {
	service ports.UserService
}

func NewHandler(service ports.UserService) *Handler {
	return &Handler{service}
}
