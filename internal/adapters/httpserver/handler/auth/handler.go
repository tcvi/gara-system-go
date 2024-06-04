package auth

import "garasystem/internal/core/ports"

type Handler struct {
	service ports.UserService
	sns     ports.SNSService
}

func NewHandler(service ports.UserService, sns ports.SNSService) *Handler {
	return &Handler{service, sns}
}
