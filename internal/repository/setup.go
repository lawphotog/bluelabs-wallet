package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// this is not needed in real app
func (r *DynamoRepository) Setup() {
	// r.deleteTable()

	var unit int64 = 1
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("userId"),
				AttributeType: "S",
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("userId"),
				KeyType:       "HASH",
			},
		},
		TableName: &tableName,
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  &unit,
			WriteCapacityUnits: &unit,
		},
	}

	//this doesn't replace existing table.
	_, err := r.client.CreateTable(context.Background(), input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Created the table", tableName)
}

func (d *DynamoRepository) deleteTable() {
	input := &dynamodb.DeleteTableInput{
		TableName: &tableName,
	}

	_, err := d.client.DeleteTable(context.Background(), input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted the table", tableName)
}
