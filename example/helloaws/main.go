package main

import (
	"log"
	"readyperfectly/aws"
)

func main() {

	awsServiceFactory, err := aws.GetAwsServiceFactory("zhixian")
	if err != nil {
		log.Fatalf("Initialization error: %v", err)
	}
	// log.Println(awsServiceFactory)

	dynamoDbService := awsServiceFactory.GetDynamoDbService()
	// log.Println(dynamoDbService)
	dynamoDbService.GetTableList()

}
