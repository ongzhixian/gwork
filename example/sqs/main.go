package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
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
	sqsClient := sqs.NewFromConfig(cfg)

	fmt.Println("Let's list the queues for your account.")
	var queueUrls []string
	paginator := sqs.NewListQueuesPaginator(sqsClient, &sqs.ListQueuesInput{})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			log.Printf("Couldn't get queues. Here's why: %v\n", err)
			break
		} else {
			queueUrls = append(queueUrls, output.QueueUrls...)
		}
	}
	if len(queueUrls) == 0 {
		fmt.Println("You don't have any queues!")
	} else {
		for _, queueUrl := range queueUrls {
			fmt.Printf("\t%v\n", queueUrl)
		}
	}
}
