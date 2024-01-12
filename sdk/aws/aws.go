package utils

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DigitalOceanSpacesStorage implements the CloudStorage interface for DigitalOcean Spaces
type DigitalOceanSpacesStorage struct {
	SpaceName string
	Region    string
	Client    *s3.S3
}

// NewDigitalOceanSpacesStorage creates a new instance of DigitalOceanSpacesStorage
func NewDigitalOceanSpacesStorage(spaceName, region, accessKeyID, secretAccessKey string) *DigitalOceanSpacesStorage {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Endpoint:    aws.String(fmt.Sprintf("https://%s.digitaloceanspaces.com", region)),
	}))

	return &DigitalOceanSpacesStorage{
		SpaceName: spaceName,
		Region:    region,
		Client:    s3.New(sess),
	}
}

// UploadFile uploads a file to DigitalOcean Spaces
func (d *DigitalOceanSpacesStorage) UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	key := fmt.Sprintf("%s/%s", folder, file.Filename)

	_, err = d.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(d.SpaceName),
		Key:    aws.String(key),
		Body:   src,
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.%s.digitaloceanspaces.com/%s", d.SpaceName, d.Region, key)
	return url, nil
}

// UploadSavedFile uploads a file from the server's local filesystem to DigitalOcean Spaces
func (d *DigitalOceanSpacesStorage) UploadSavedFile(filePath string, folder string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	key := fmt.Sprintf("%s/%s", folder, file.Name())

	_, err = d.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(d.SpaceName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.%s.digitaloceanspaces.com/%s", d.SpaceName, d.Region, key)
	return url, nil
}

// DeleteFile deletes a file from DigitalOcean Spaces
func (d *DigitalOceanSpacesStorage) DeleteFile(filePath string) error {
	_, err := d.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(d.SpaceName),
		Key:    aws.String(filePath),
	})
	return err
}
