package aws

import (
	"context"
	"fmt"
	"readyperfectly/aws/dynamoDb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsServiceFactory struct {
	ProfileName string
	AwsConfig   aws.Config
	Context     context.Context
}

func GetAwsServiceFactory(profileName string) (*AwsServiceFactory, error) {
	backgroundContext := context.Background()
	cfg, err := config.LoadDefaultConfig(
		backgroundContext,
		config.WithSharedConfigProfile(profileName),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config for profile %s: %w", profileName, err)
	}

	return &AwsServiceFactory{
		ProfileName: profileName,
		AwsConfig:   cfg,
		Context:     backgroundContext,
	}, nil
}

func (factory AwsServiceFactory) GetDynamoDbService() *dynamoDb.DynamoDbService {
	return dynamoDb.NewDynamoDbService(factory.Context, factory.AwsConfig)
}
