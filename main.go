package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Credentials.
	accessKey := "AWS_ACCESS_KEY_ID"
	secretKey := "AWS_SECRET_ACCESS_KEY"
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

	// S3 Client.
	config := aws.Config{
		Credentials: creds,
		Region:      "ap-northeast-1",
	}
	client := s3.New(&config)

	// Open file.
	file, err := os.Open("FILENAME")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read file.
	byteArray, err := ioutil.ReadAll(file)

	params := &s3.PutObjectInput{
		Bucket: aws.String("S3_BUCKET_NAME"),
		Key:    aws.String("S3_KEY"),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(byteArray),
	}

	// Put object.
	resp, err := client.PutObject(params)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
