package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

func main() {
	ctx := context.Background()

	// Using the SDK's default configuration, load additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	//cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("zhixian"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the SQS client
	snsClient := sns.NewFromConfig(cfg)
	fmt.Println("Let's list the topics for your account.")
	var topics []types.Topic
	paginator := sns.NewListTopicsPaginator(snsClient, &sns.ListTopicsInput{})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			log.Printf("Couldn't get topics. Here's why: %v\n", err)
			break
		} else {
			topics = append(topics, output.Topics...)
		}
	}
	if len(topics) == 0 {
		fmt.Println("You don't have any topics!")
	} else {
		for _, topic := range topics {
			fmt.Printf("\t%v\n", *topic.TopicArn)
		}
	}
}
