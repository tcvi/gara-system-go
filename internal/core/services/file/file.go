package file

import (
	"garasystem/internal/core/ports"
	"garasystem/internal/core/services"
)

type Service struct {
	repo *services.Repository
	s3   ports.FileStorage
}

func NewService(repo *services.Repository, s3Storage ports.FileStorage) *Service {
	return &Service{repo: repo, s3: s3Storage}
}
