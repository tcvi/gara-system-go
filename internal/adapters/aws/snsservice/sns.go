package snsservice

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SnsService struct {
	client *sns.Client
}

func NewSnsService(cfg aws.Config) *SnsService {
	svc := sns.NewFromConfig(cfg)
	return &SnsService{client: svc}
}
