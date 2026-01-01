package devops

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getAWSRegion() string {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-west-2"
	}
	return region
}

func getSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(getAWSRegion())}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return sess
}

func getBucketName() string {
	return os.Getenv("BUCKET_NAME")
}

func getBucket(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

func removePrefixFromPaths(s string) string {
	parts := strings.Split(s, "/")
	if len(parts) > 1 {
		return strings.Join(parts[1:], "/")
	}
	return s
}

func getLocalDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func getLocalFileContent(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func getRemoteFileContent(bucket *s3.S3, key string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(getBucketName()),
		Key:    aws.String(key),
	}
	ret, err := bucket.GetObject(context.TODO(), input)
	if err != nil {
		return "", err
	}
	return string(ret.Body), nil
}