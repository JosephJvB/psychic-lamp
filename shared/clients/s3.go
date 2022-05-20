package clients

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func init() {
	log.SetPrefix("clients.s3: ")
	log.SetFlags(0)
	log.Println("init()")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	
	s3Client = s3.NewFromConfig(cfg)
}
func S3GetObject(bucket string, key string) (*s3.GetObjectOutput, error) {
	log.Println("GetObject()", bucket, key)
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	}
	return s3Client.GetObject(context.TODO(), params)
}
