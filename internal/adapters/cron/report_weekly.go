package cron

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/logger"
)

func (s *Server) ReportWeekly() {
	logger.Log.Println("ReportWeekly")
	err := s.notificationService.PushNotifications(
		[]string{}, domain.Notification{
			Title:   "Report weekly",
			Message: "This is auto message",
		},
	)

	if err != nil {
		logger.Log.Errorln("failed push notification weekly: ", err)
	}
}
