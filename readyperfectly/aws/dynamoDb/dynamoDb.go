package dynamoDb

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDbService struct {
	Context        context.Context
	DynamoDbClient *dynamodb.Client
}

func NewDynamoDbService(context context.Context, awsConfig aws.Config) *DynamoDbService {
	return &DynamoDbService{
		Context:        context,
		DynamoDbClient: dynamodb.NewFromConfig(awsConfig),
	}
}

func (service *DynamoDbService) GetTableList() {
	resp, err := service.DynamoDbClient.ListTables(service.Context, &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range resp.TableNames {
		fmt.Println(tableName)
	}
}
