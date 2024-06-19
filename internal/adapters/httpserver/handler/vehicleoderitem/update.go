package vehicleoderitem

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateOrderItem(c echo.Context) error {
	var req domain.UpdateVehicleOrderItemRequest
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrVehicleOrderItemInvalidData(err))
	}

	err = h.service.Updates(req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
