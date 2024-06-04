package user

import "garasystem/internal/core/domain"

func (u *Service) UpdateUser(id, email, password string) error {
	return u.userStorage().Update(&domain.User{})
}
