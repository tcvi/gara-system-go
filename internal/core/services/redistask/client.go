package redistask

import (
	"encoding/json"
	"fmt"
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/config"
	"github.com/hibiken/asynq"
)

type Client struct {
	clientTask *asynq.Client
}

func NewRedisTaskClient(config *config.Config) *Client {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     fmt.Sprint(config.Redis.Host, ":", config.Redis.Port),
		Password: config.Redis.Password,
	})
	return &Client{client}
}

func (c *Client) NewPushNotificationTask(payload domain.TaskPushNotificationPayload) error {
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return myerror.ErrRedisTaskMarshalPushNotificationPayload(err)
	}

	task := asynq.NewTask(domain.TypePushNotification, payloadByte)

	_, err = c.clientTask.Enqueue(task)
	if err != nil {
		return myerror.ErrRedisTaskEnqueue(err)
	}

	return nil
}
