package user

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Storage struct {
	db *dynamodb.Client
}

func NewStorage(db *dynamodb.Client) *Storage {
	return &Storage{db}
}
