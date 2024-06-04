package user

import (
	"github.com/pkg/errors"
)

var (
	userNotFoundError = errors.New("User not found")
)

func (s *Storage) IsExisted(userName string, email string, phone string) error {
	// Check by userName
	if userName != "" {
		_, err := s.Get("user_name = ?", userName)
		if !errors.Is(err, userNotFoundError) {
			return errors.New("Existed username")
		}
	}

	// Check by phoneNumber
	if phone != "" {
		_, err := s.Get("phone_number = ?", phone)
		if !errors.Is(err, userNotFoundError) {
			return errors.New("Existed phone")
		}
	}

	// Check by email
	if email != "" {
		_, err := s.Get("email = ?", email)
		if !errors.Is(err, userNotFoundError) {
			return errors.New("Existed email")
		}
	}

	return nil
}
