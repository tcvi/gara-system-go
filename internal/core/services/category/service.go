package category

import (
	"garasystem/internal/core/services"
)

type Service struct {
	repo *services.Repository
}

func NewService(repo *services.Repository) *Service {
	return &Service{repo: repo}
}
