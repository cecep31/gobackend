package storage

import (
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewFileStorageClient() (*minio.Client, error) {
	endpoint := os.Getenv("S3_ENDPOINT")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")

	creds := credentials.NewStaticV4(accessKey, secretKey, "")
	client, err := minio.New(endpoint, &minio.Options{Creds: creds, Secure: true})
	if err != nil {
		return nil, err
	}

	return client, nil
}
