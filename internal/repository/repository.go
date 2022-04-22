package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoRepository struct {
	client *dynamodb.Client
}

type Repository interface {
	Create(userId string) error
}

func New(client *dynamodb.Client) *DynamoRepository {
	return &DynamoRepository{
		client: client,
	}
}

var tableName string = "wallet"

func (r *DynamoRepository) Create(userId string) error {
	item := &Wallet{
		UserId: userId,
		UpdateSequence: "0",
	}

	data, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("MarshalMap: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      data,
		ConditionExpression: aws.String("attribute_not_exists(userId)"), //won't replace if exists
	}

	_, err = r.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("PutItem: %v", err)
	}

	return nil
}