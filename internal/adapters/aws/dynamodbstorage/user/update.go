package user

import (
	"context"
	"fmt"
	"garasystem/internal/core/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"
)

func (s *Storage) Update(user *domain.User) error {
	user.UpdatedAt = time.Now()

	update := expression.Set(expression.Name("user_name"), expression.Value(user.UserName))
	update.Set(expression.Name("password"), expression.Value(user.Password))
	update.Set(expression.Name("phone_number"), expression.Value(user.PhoneNumber))
	update.Set(expression.Name("is_active"), expression.Value(user.IsActive))
	update.Set(expression.Name("active_code"), expression.Value(user.ActiveCode))
	update.Set(expression.Name("exp_code"), expression.Value(user.ExpCode))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return nil
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: fmt.Sprint(user.ID)},
		},
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	}

	_, err = s.db.UpdateItem(context.TODO(), input)
	return err
}
