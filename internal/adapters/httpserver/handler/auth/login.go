package auth

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	var req domain.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrInvalidRegister(err))
	}

	res, err := h.service.Login(req, h.config.JwtSecretKey)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, res)
}
