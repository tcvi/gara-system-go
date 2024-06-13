package category

import (
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
	"strconv"
)

func (h *Handler) Delete(c echo.Context) error {
	categoryIDStr := c.Param("id")

	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		return util.Response.Error(c, myerror.ErrCategoryInvalidParamID(err))
	}

	err = h.service.Delete(categoryID)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
