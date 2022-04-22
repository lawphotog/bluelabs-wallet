package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoRepository struct {
	client *dynamodb.Client
}

type Repository interface {
	Create(userId string) error
	Get(userId string) (Wallet, error)
	Update(wallet Wallet) error
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

func (r *DynamoRepository) Update(wallet Wallet) error{
	data, err := attributevalue.MarshalMap(wallet)
	if err != nil {
		return fmt.Errorf("MarshalMap: %v", err)
	}

	updateSeq, _ := strconv.Atoi(wallet.UpdateSequence)
	sequence := strconv.Itoa(updateSeq - 1)

	input := &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      data,
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":seq": &types.AttributeValueMemberS{Value: sequence},
		},
		ConditionExpression: aws.String("UpdateSequence = :seq"),
	}

	_, err = r.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("PutItem: %v", err)
	}

	return nil
}

func (r *DynamoRepository) Get(userId string) (Wallet, error) {
	items := []Wallet{}
	input := &dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("userId = :userId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userId": &types.AttributeValueMemberS{Value: userId},
		},
	}

	data, err := r.client.Query(context.TODO(), input)
	if err != nil {
		fmt.Println("error")
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &items)
	if err != nil {
		fmt.Printf("UnmarshalListOfMaps: %v", err)
	}

	if len(items) < 1 {
		fmt.Println("error")
		return Wallet{}, errors.New("not found")
	}

	return items[0], nil
}
