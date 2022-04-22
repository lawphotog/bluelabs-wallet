package repository

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoRepository struct {
	client *dynamodb.Client
}

type Repository interface {
}

func New(client *dynamodb.Client) *DynamoRepository {
	return &DynamoRepository{
		client: client,
	}
}