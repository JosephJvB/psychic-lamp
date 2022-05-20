package main

import (
	"fmt"
	"log"
	"os"

	"multi-lambda/shared/clients"
	"multi-lambda/shared/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.SetPrefix("Lambda1:")
	log.SetFlags(0)
}

func HandleLambdaEvent(event models.MyRequest) (models.MyResponse, error) {
	log.Println("env.Debug:", os.Getenv("DEBUG"))
	log.Println("event:", event)
	log.Println("event.Data:", event.Data)
	bucket := "hubspot-csv-backup"
	key := "descriptions.json"
	responseData := "success"
	r, err := clients.S3GetObject(bucket, key)
	if err != nil {
		log.Fatal(err)
		responseData = fmt.Sprintf("failed to get object s3://%s/%s", bucket, key)
	}
	responseData = fmt.Sprintf("object ContentLength: %d", r.ContentLength)
	return models.MyResponse {
		Data: responseData,
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}