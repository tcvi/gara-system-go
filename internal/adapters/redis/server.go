package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"garasystem/internal/core/domain"
	"garasystem/internal/core/ports"
	"garasystem/internal/logger"
	"garasystem/pkg/config"
	"github.com/hibiken/asynq"
)

type Server struct {
	TaskHandler ports.RedisTaskHandler
}

func (s *Server) HandlePushNotification(ctx context.Context, task *asynq.Task) error {
	var p domain.TaskPushNotificationPayload
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	return s.TaskHandler.HandlePushNotification(p)
}

func NewServer(config *config.Config, notificationService ports.NotificationService) {
	s := Server{
		TaskHandler: &Handler{
			notificationService,
		},
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     fmt.Sprint(config.Redis.Host, ":", config.Redis.Port),
			Password: config.Redis.Password,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(domain.TypePushNotification, s.HandlePushNotification)
	if err := srv.Run(mux); err != nil {
		logger.Log.Fatalf("could not run server: %v", err)
	}
}
