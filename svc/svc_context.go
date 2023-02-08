package svc

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/szpnygo/goclip/clip"
	"github.com/szpnygo/goclip/weaviatex"
)

type ServiceContext struct {
	S3Client   *s3.S3
	Bucket     string
	ClipHelper *clip.ClipHelper
}

func NewServiceContext() *ServiceContext {
	session, err := newSession()
	if err != nil {
		panic(err)
	}
	client := s3.New(session)

	bucket := os.Getenv("S3_BUCKET")
	if len(bucket) == 0 {
		panic("S3_BUCKET is empty")
	}

	weaviateHost := os.Getenv("WEAVIATE_HOST")
	if len(weaviateHost) == 0 {
		panic("WEAVIATE_HOST is empty")
	}
	weaviateClient, err := weaviatex.NewWeaviateClient(weaviateHost)
	if err != nil {
		panic(err)
	}
	clipHelper := clip.NewClipHelper(weaviateClient)
	err = clipHelper.Init("GoClip")
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		S3Client:   client,
		Bucket:     bucket,
		ClipHelper: clipHelper,
	}
}

func newSession() (*session.Session, error) {
	secretID := os.Getenv("S3_ID")
	secretKEY := os.Getenv("S3_KEY")
	creds := credentials.NewStaticCredentials(secretID, secretKEY, "")
	region := os.Getenv("S3_REGION")
	endpoint := os.Getenv("S3_ENDPOINT")

	config := &aws.Config{
		Region:           aws.String(region),
		Endpoint:         &endpoint,
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
	}

	return session.NewSession(config)
}
