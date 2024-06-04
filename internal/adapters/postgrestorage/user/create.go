package user

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Create(user *domain.User) error {
	err := s.db.Create(user).Error
	return errors.Wrap(err, "Create user fail")
}
