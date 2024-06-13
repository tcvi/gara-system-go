package user

import (
	"garasystem/internal/core/domain"
	"github.com/pkg/errors"
)

func (s *Storage) GetList(query interface{}, args ...interface{}) ([]domain.User, error) {
	var users []domain.User

	err := s.db.Where(query, args...).Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, "GetList User fail")
	}

	return users, nil
}
