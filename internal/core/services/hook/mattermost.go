package hook

import (
	"fmt"
	"garasystem/internal/core/domain"
	"garasystem/pkg/config"
	"github.com/go-resty/resty/v2"
)

type MattermostHook struct {
	host string
}

func NewMattermostHook(config *config.Config) *MattermostHook {
	return &MattermostHook{
		host: config.Hook.Mattermost,
	}
}

func (m *MattermostHook) Send(mess string) error {
	if m.host == "" {
		return nil
	}

	client := resty.New()

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(domain.MattermostPostReq{Text: mess}).
		Post(m.host)
	if err != nil {
		fmt.Print(err)
	}
	return nil
}
