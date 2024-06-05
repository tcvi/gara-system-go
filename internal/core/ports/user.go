package ports

import (
	"garasystem/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type UserStore interface {
	Get(query interface{}, args ...interface{}) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	IsExisted(userName string, email string, phone string) error
}

type UserHandler interface {
	Update(c echo.Context) error
}

type UserService interface {
	CreateUser(req domain.RegisterUserReq) error
	UpdateUser(id, email, password string) error
	VerifyAccount(req domain.VerifyRequest) error
	ResendCode(req domain.ResendCodeRequest) error
	Login(req domain.LoginRequest, secretKey string) (*domain.UserLoginResponse, error)
}
