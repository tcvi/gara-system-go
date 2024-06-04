package user

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) Update(user *domain.User) error {
	err := s.db.Updates(user).Error
	return errors.Wrap(err, "Update user fail")
}
