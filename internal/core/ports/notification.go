package ports

import (
	"garasystem/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type NotificationService interface {
	PushNotifications(tokens []string, req domain.Notification) error
}

type NotificationHandler interface {
	PushNotification(c echo.Context) error
}
