package auth

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ResendCode(c echo.Context) error {
	var req domain.ResendCodeRequest
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrInvalidRegister(err))
	}

	err = h.service.ResendCode(req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
