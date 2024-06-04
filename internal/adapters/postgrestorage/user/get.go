package user

import (
	"garasystem/internal/core/domain"
)

func (s *Storage) Get(query interface{}, args ...interface{}) (*domain.User, error) {
	var user domain.User

	err := s.db.Where(query, args).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
