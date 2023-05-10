package storage

import (
	"log"
	"os"

	"github.com/gofiber/storage/s3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	Storage     *s3.Storage
	FileStorage *minio.Client
)

func IniStorage() {
	Storage = s3.New(s3.Config{
		Bucket: os.Getenv("S3_BUCKET"),
		Region: os.Getenv("S3_REGION"),
		Reset:  false,
		Credentials: s3.Credentials{
			AccessKey:       os.Getenv("S3_ACCESS_KEY"),
			SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
		},
	})
}

func InitFileStorage() {
	minioClient, err := minio.New(os.Getenv("S3_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("S3_ACCESS_KEY"), os.Getenv("S3_SECRET_KEY"), ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	FileStorage = minioClient

}
