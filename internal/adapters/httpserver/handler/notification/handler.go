package notification

import "garasystem/internal/core/ports"

type Handler struct {
	service ports.NotificationService
}

func NewHandler(service ports.NotificationService) *Handler {
	return &Handler{service}
}
