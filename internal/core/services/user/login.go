package user

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u *Service) validateLogin(req domain.LoginRequest) error {
	if req.UserName == "" || req.Password == "" {
		return myerror.ErrInvalidLogin(nil)
	}

	return nil
}

func (u *Service) Login(req domain.LoginRequest, secretKey string) (*domain.UserLoginResponse, error) {
	if err := u.validateLogin(req); err != nil {
		return nil, err
	}

	user, err := u.userStorage().Get("user_name = ?", req.UserName)
	if err != nil {
		return nil, myerror.ErrUserNotFound(err)
	}

	if !user.IsActive {
		return nil, myerror.ErrAuthAccountNotActive(nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, myerror.ErrAuthWrongPassword(err)
	}

	res := user.MappingUserLogin()

	token, err := jwt.GenerateToken(res.ID, secretKey)
	if err != nil {
		return nil, myerror.ErrAuthGenerateToken(err)
	}

	res.Token = token

	return res, nil
}
