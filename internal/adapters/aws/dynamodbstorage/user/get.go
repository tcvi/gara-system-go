package user

import (
	"context"
	"fmt"
	"garasystem/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	"strings"
)

func (s *Storage) Get(query interface{}, args ...interface{}) (*domain.User, error) {
	value := args[0].(string)

	if strings.Index(query.(string), "user_name") != -1 {
		return getUserByAttribute(s.db, "UserNameIndex", "user_name", value)
	}

	if strings.Index(query.(string), "phone_number") != -1 {
		return getUserByAttribute(s.db, "PhoneNumberIndex", "phone_number", value)
	}

	if strings.Index(query.(string), "email") != -1 {
		return getUserByAttribute(s.db, "EmailIndex", "email", value)
	}

	return nil, errors.New("Not implement")
}

func getUserByAttribute(svc *dynamodb.Client, indexName string, attributeName string, attributeValue string) (*domain.User, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("Users"),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :v", attributeName)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": &types.AttributeValueMemberS{Value: attributeValue},
		},
		Limit: aws.Int32(1),
	}

	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Count == 0 {
		return nil, userNotFoundError
	}

	var user domain.User
	err = attributevalue.UnmarshalMap(result.Items[0], &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
