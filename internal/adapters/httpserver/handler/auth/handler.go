package auth

import (
	"garasystem/internal/core/ports"
	"garasystem/pkg/config"
)

type Handler struct {
	config  *config.Config
	service ports.UserService
	sns     ports.SNSService
}

func NewHandler(config *config.Config, service ports.UserService, sns ports.SNSService) *Handler {
	return &Handler{config, service, sns}
}
