package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Service struct {
	client *s3.Client
}

func (s *Service) PresignClient() *s3.PresignClient {
	return s3.NewPresignClient(s.client)
}

func NewS3Service(cfg aws.Config) *Service {
	s3Client := s3.NewFromConfig(cfg)
	return &Service{s3Client}
}
