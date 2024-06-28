package domain

// A list of task types.
const (
	TypePushNotification = "push:notification"
)

type TaskPushNotificationPayload struct {
	Tokens       []string
	Notification Notification
}
