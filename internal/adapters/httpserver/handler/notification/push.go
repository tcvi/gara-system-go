package notification

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PushNotification(c echo.Context) error {
	var req domain.PushNotificationReq
	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(c, myerror.ErrNotificationInvalidData(err))
	}

	err = h.service.PushNotifications(req.Tokens, req.Data)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, nil)
}
