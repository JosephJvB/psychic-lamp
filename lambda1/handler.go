package main

import (
	"fmt"
	"log"
	"strconv"

	"multi-lambda/shared/clients"
	"multi-lambda/shared/models"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	log.SetPrefix("Lambda1:")
	log.SetFlags(0)
	// do startup work here, init clients etc
}

func HandleLambdaEvent(event models.MyRequest) (models.MyResponse, error) {
	fmt.Println("--- event ---")
	fmt.Println(event)
	fmt.Println("--- event.Data ---")
	fmt.Println(event.Data)
	responseData := "success"
	r, err := clients.S3GetObject("hubspot-csv-backup", "descriptions.json")
	if err != nil {
		log.Fatal(err)
		responseData = "failed to get object"
	}
	responseData = "object ContentLength: " + strconv.FormatInt(r.ContentLength, 10)
	return models.MyResponse {
		Data: responseData,
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}