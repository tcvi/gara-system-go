package s3

import (
	"context"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

func (s *Service) PresignObject(objectKey string, expire time.Duration) (string, error) {
	cfg := config.GetConfig()

	request, err := s.PresignClient().PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(cfg.AWS.BucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expire
	})
	if err != nil {
		return "", myerror.ErrFilePresignS3(err)
	}

	return request.URL, nil
}
