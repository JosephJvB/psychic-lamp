package main

import (
	"log"

	"multi-lambda/shared/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.SetPrefix("Lambda2:")
	log.SetFlags(0)
}

func HandleLambdaEvent(event models.MyRequest) (models.MyResponse, error) {
	log.Println("event:", event)
	log.Println("event.Data:", event.Data)
	return models.MyResponse {
		Data: "yo",
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}