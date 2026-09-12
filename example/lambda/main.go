package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
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

	// Using the Config value, create the Lambda client
	lambdaClient := lambda.NewFromConfig(cfg)

	maxItems := 10
	fmt.Printf("Let's list up to %v functions for your account.\n", maxItems)
	result, err := lambdaClient.ListFunctions(ctx, &lambda.ListFunctionsInput{
		MaxItems: aws.Int32(int32(maxItems)),
	})
	if err != nil {
		fmt.Printf("Couldn't list functions for your account. Here's why: %v\n", err)
		return
	}
	if len(result.Functions) == 0 {
		fmt.Println("You don't have any functions!")
	} else {
		for _, function := range result.Functions {
			fmt.Printf("\t%v\n", *function.FunctionName)
		}
	}
}
