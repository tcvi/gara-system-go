package user

import (
	"context"
	"garasystem/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
	"time"
)

func (s *Storage) Create(user *domain.User) error {
	now := time.Now()

	user.ID = now.Unix()
	user.CreatedAt = now
	user.UpdatedAt = now

	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return errors.Wrap(err, "")
	}

	_, err = s.db.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{TableName: aws.String("Users"), Item: item},
	)

	return errors.Wrap(err, "Create user fail")
}
