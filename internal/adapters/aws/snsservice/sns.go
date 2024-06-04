package snsservice

import (
	"context"
	"garasystem/internal/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SnsService struct {
	client *sns.Client
}

func NewSnsService(cfg aws.Config) *SnsService {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-southeast-1"),
		config.WithSharedConfigProfile("gara-system-sns-dev"),
	)
	if err != nil {
		logger.Log.Fatalln(err)
	}

	svc := sns.NewFromConfig(cfg)
	return &SnsService{client: svc}
}
