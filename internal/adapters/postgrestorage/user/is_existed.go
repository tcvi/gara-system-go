package user

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Storage) IsExisted(userName string, email string, phone string) error {
	if userName != "" {
		user, err := s.Get("user_name = ?", userName)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if user != nil {
			return errors.New("Username existed")
		}
	}

	if email != "" {
		user, err := s.Get("email = ?", email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if user != nil {
			return errors.New("Email existed")
		}
	}

	if phone != "" {
		user, err := s.Get("phone_number = ?", phone)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if user != nil {
			return errors.New("Phone existed")
		}
	}

	return nil
}
