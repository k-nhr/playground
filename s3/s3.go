package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	file, err := os.Open("test01.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	creds := credentials.NewStaticCredentials("", "", "")

	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)
	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(sess)
	out, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("test-2019-04-16"),
		Key:    aws.String("test01.txt"),
		Body:   file,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
	return
}
