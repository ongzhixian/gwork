package dynamoDb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDbService struct {
	DynamoDbClient *dynamodb.Client
}

func NewDynamoDbService(awsConfig aws.Config) *DynamoDbService {
	return &DynamoDbService{
		DynamoDbClient: dynamodb.NewFromConfig(awsConfig),
	}
}
