package vehicleorder

import (
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
	"strconv"
)

func (h *Handler) GetByID(c echo.Context) error {
	idStr := c.Param("id")

	recordID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.Response.Error(c, myerror.ErrVehicleInvalidIDParam(err))
	}

	vehicle, err := h.service.GetByID(recordID)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, vehicle)
}
