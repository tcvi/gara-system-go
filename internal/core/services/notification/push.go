package notification

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
)

func (s *Service) PushNotifications(tokens []string, req domain.Notification) error {
	a, err := s.Client.Messaging(context.Background())
	if err != nil {
		return myerror.ErrNotificationClientMessage(err)
	}

	_, err = a.SendEachForMulticast(context.Background(), &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: req.Title,
			Body:  req.Message,
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Data: map[string]string{
				"sound": "default",
			},
		},
	})
	if err != nil {
		return myerror.ErrNotificationPush(err)
	}

	return nil
}
