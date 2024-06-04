package dynamodbstorage

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewConnection(cfg aws.Config) (*dynamodb.Client, error) {
	return dynamodb.NewFromConfig(cfg), nil
}
