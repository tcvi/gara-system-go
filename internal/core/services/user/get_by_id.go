package user

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (u *Service) GetByID(id int64) (*domain.UserModel, error) {
	user, err := u.userStorage().Get("id = ?", id)
	if err != nil {
		return nil, myerror.ErrUserNotFound(err)
	}

	if !user.IsActive {
		return nil, myerror.ErrAuthAccountNotActive(nil)
	}

	return user.MappingUserModel(), nil
}
