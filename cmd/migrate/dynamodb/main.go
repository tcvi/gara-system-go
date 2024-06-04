package main

import (
	"context"
	"garasystem/internal/adapters/aws/dynamodbstorage"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

func main() {
	db, err := dynamodbstorage.NewConnection()
	if err != nil {
		log.Fatalln("Connect dynamodb fail: ", err)
	}

	_, err = createUserTable(db)
	if err != nil {
		log.Fatalln("Create Users table fail: ", err)
	}
	log.Println("Create Users table success")
}

func createUserTable(db *dynamodb.Client) (*types.TableDescription, error) {
	tableName := "Users"
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeN,
			},
			{
				AttributeName: aws.String("user_name"),
				AttributeType: types.ScalarAttributeTypeS, // String
			},
			{
				AttributeName: aws.String("phone_number"),
				AttributeType: types.ScalarAttributeTypeS, // String
			},
			{
				AttributeName: aws.String("email"),
				AttributeType: types.ScalarAttributeTypeS, // String
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName: aws.String(tableName),
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("UserNameIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("user_name"),
						KeyType:       types.KeyTypeHash, // Partition key
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
			{
				IndexName: aws.String("PhoneNumberIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("phone_number"),
						KeyType:       types.KeyTypeHash, // Partition key
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
			{
				IndexName: aws.String("EmailIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("email"),
						KeyType:       types.KeyTypeHash, // Partition key
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}

	table, err := db.CreateTable(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return table.TableDescription, nil
}
