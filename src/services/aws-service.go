package services

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	Client     *s3.Client
	BucketName string
	Region     string
}

func NewS3Uploader(bucketName string, region string) (S3Uploader, error) {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return S3Uploader{}, fmt.Errorf("unable to load SDK config, %v", err)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	return S3Uploader{
		Client:     client,
		BucketName: bucketName,
		Region:     region,
	}, nil
}

func (u *S3Uploader) UploadImage(file multipart.File, fileName string) (string, error) {
	// Upload the file to S3
	_, err := u.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(u.BucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.BucketName, u.Region, fileName)

	return fileURL, nil
}
