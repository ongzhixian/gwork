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

	// Access factory.Config to construct S3, DynamoDB, or other AWS clients

	// awsServiceFactory := aws.AwsServiceFactory{
	// 	ProfileName: "zhixian",
	// }
	log.Println(awsServiceFactory)

	dynamoDbService := awsServiceFactory.GetDynamoDbService()
	log.Println(dynamoDbService)

	// dynamoDbService := dynamoDb.DynamoDbService{
	// 	Profile: "SomeProf",
	// }

	// fmt.Print(dynamoDbService)

	// dynamoDb.DoDynamo()

	// asd := DynamoDbService{
	// 	Profile: "asd",
	// }

	// dynamoDbService.SomePackageFunction()

}
