package ports

import (
	"garasystem/internal/core/domain"
)

type RedisTaskClient interface {
	NewPushNotificationTask(payload domain.TaskPushNotificationPayload) error
}

type RedisTaskHandler interface {
	HandlePushNotification(payload domain.TaskPushNotificationPayload) error
}
