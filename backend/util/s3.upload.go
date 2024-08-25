package util

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetS3Client retorna um cliente S3 configurado com as credenciais e região
func GetS3Client() (*s3.Client, error) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		return nil, fmt.Errorf("missing AWS_REGION in environment variables")
	}

	keyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if keyID == "" || secretKey == "" {
		return nil, fmt.Errorf("missing AWS_ACCESS_KEY_ID or AWS_SECRET_ACCESS_KEY in environment variables")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(keyID, secretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	return s3Client, nil
}

func GetS3Bucket() (string, error) {
	bucketName := os.Getenv("S3_BUCKET_NAME")
	if bucketName == "" {
		return "", fmt.Errorf("missing S3_BUCKET_NAME in environment variables")
	}
	return bucketName, nil
}

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	s3Client, err := GetS3Client()
	if err != nil {
		return "", err
	}

	bucketName, err := GetS3Bucket()
	if err != nil {
		return "", err
	}

	uploader := manager.NewUploader(s3Client)

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	timestamp := time.Now().Format("15040502012006")
	newFileName := fmt.Sprintf("%s_%s", timestamp, fileHeader.Filename)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(newFileName),
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: aws.String(http.DetectContentType(buffer.Bytes())),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// Construindo a URL do arquivo no S3
	region := os.Getenv("AWS_REGION")
	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, newFileName)
	return fileURL, nil
}

func ListObjectsInBucket() ([]string, error) {
	s3Client, err := GetS3Client()
	if err != nil {
		return nil, err
	}

	bucketName, err := GetS3Bucket()
	if err != nil {
		return nil, err
	}

	var objectKeys []string
	paginator := s3.NewListObjectsV2Paginator(s3Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %w", err)
		}

		for _, object := range output.Contents {
			objectKeys = append(objectKeys, *object.Key)
		}
	}

	return objectKeys, nil
}

func DeleteObjectFromS3(key string) error {
	s3Client, err := GetS3Client()
	if err != nil {
		return err
	}

	bucketName, err := GetS3Bucket()
	if err != nil {
		return err
	}

	_, err = s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object from S3: %w", err)
	}

	return nil
}
