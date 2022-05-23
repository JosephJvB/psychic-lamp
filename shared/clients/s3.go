package clients

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type IS3 interface {
	GetObject(bucket string, key string)
}
type S3 struct {
	client *s3.Client
}

func NewS3Client() *S3 {
	log.SetPrefix("clients.s3: ")
	log.SetFlags(0)
	log.Println("init()")
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	_client := s3.NewFromConfig(cfg)
	c := &S3{ _client }
	return c
}
func (s *S3) GetObject(bucket string, key string) (*s3.GetObjectOutput, error) {
	log.Println("GetObject()", bucket, key)
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	}
	return s.client.GetObject(context.TODO(), params)
}
