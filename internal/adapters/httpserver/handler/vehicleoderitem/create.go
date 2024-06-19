package vehicleoderitem

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateOrderItem(c echo.Context) error {
	var req domain.CreateVehicleOrderItemRequest
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrVehicleOrderItemInvalidData(err))
	}

	err = h.service.Create(req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
