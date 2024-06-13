package category

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Update(c echo.Context) error {
	var req domain.UpdateCategoryReq
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrCategoryDataInvalid(err))
	}

	err = h.service.Update(req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
