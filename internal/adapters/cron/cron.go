package cron

import (
	"garasystem/internal/core/ports"
	"garasystem/internal/logger"
	"github.com/robfig/cron/v3"
)

var (
	EveryMondayPatter = "30 8 ? * MON" // At 08:30 AM, only on Monday
)

type Server struct {
	cron                *cron.Cron
	notificationService ports.NotificationService
}

func StartCron() *Server {
	c := cron.New()

	c.Start()
	logger.Log.Println("Start Cron")

	s := &Server{cron: c}

	_, err := c.AddFunc(EveryMondayPatter, s.ReportWeekly)
	if err != nil {
		logger.Log.Fatal("failed add ReportWeekly to CronJob: ", err)
	}

	return s
}

func (s *Server) Stop() {
	s.cron.Stop()
}
