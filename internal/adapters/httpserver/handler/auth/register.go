package auth

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context) error {
	var req domain.RegisterUserReq
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrInvalidRegister(err))
	}

	err = h.service.CreateUser(req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
