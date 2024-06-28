package redis

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/ports"
)

type Handler struct {
	NotificationService ports.NotificationService
}

func (h *Handler) HandlePushNotification(payload domain.TaskPushNotificationPayload) error {
	return h.NotificationService.PushNotifications(payload.Tokens, payload.Notification)
}
