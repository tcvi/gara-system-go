package user

import (
	"garasystem/internal/core/ports"
	"garasystem/internal/core/services"
)

type Service struct {
	repo       *services.Repository
	awsService ports.SNSService
}

func NewUserService(repo *services.Repository, awsService ports.SNSService) *Service {
	return &Service{repo: repo, awsService: awsService}
}

func (u *Service) userStorage() ports.UserStore {
	return u.repo.UserStore
}
