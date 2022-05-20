package main

import (
	"fmt"

	"multi-lambda/shared/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	// do startup work here, init clients etc
}

func HandleLambdaEvent(event models.MyRequest) (models.MyResponse, error) {
	fmt.Println("--- event ---")
	fmt.Println(event)
	fmt.Println("--- event.Data ---")
	fmt.Println(event.Data)
	return models.MyResponse {
		Data: "yo",
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}