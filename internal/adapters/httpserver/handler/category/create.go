package category

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	var req domain.CreateCategoryReq
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrCategoryDataInvalid(err))
	}

	err = h.service.Create(req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
