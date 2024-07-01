package redis

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/ports"
)

type Handler struct {
	notificationService ports.NotificationService
	hook                ports.HookService
}

func (h *Handler) HandlePushNotification(payload domain.TaskPushNotificationPayload) error {
	// Send message to mattermost
	_ = h.hook.Send(payload.Notification.Message)

	return h.notificationService.PushNotifications(payload.Tokens, payload.Notification)
}
