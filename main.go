package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)


func HandleRequest(ctx context.Context, request events.ALBTargetGroupRequest) (string, error) {
	/*
	
	log.Fatal(request.Body)
	for key, value := range request.Headers {
		log.Fatal(key, value)
	}
	*/

	
	
	creds := credentials.NewStaticCredentialsProvider("", "", "")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),  config.WithCredentialsProvider(creds), config.WithRegion("ap-northeast-2"),
	)

	if err != nil {
		log.Fatal(err)
	}

	
	svc := s3.NewFromConfig(cfg)

	presignClient := s3.NewPresignClient(svc)

	fmt.Print(presignClient)

	

	presignedUrl, err := presignClient.PresignPutObject(context.Background(),
	&s3.PutObjectInput{
		Bucket: aws.String(""),
		Key: aws.String("memberkey"),
	},
	s3.WithPresignExpires(time.Minute*15))

	if err != nil {
		log.Fatal(err)
		return "", err;
	}

	url := presignedUrl.URL

	fmt.Print(url)

	return url, nil	
}


func main() {
	lambda.Start(HandleRequest)
}