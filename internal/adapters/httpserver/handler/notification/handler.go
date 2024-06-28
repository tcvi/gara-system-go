package notification

import "garasystem/internal/core/ports"

type Handler struct {
	service   ports.NotificationService
	RedisTask ports.RedisTaskClient
}

func NewHandler(service ports.NotificationService, redisTaskClient ports.RedisTaskClient) *Handler {
	return &Handler{service, redisTaskClient}
}
