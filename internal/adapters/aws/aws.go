package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig() (cfg aws.Config, err error) {
	return config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-southeast-1"),
		config.WithSharedConfigProfile("gara-system-dev"),
	)
}
