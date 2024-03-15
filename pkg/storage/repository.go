package storage

import (
	"mime/multipart"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/valyala/fasthttp"
)

type Repository interface {
	UploadFile(ctx *fasthttp.RequestCtx, objectname string, file *multipart.FileHeader) error
}

type repository struct {
	minio *minio.Client
}

func NewRepo(minioclient *minio.Client) Repository {
	return &repository{
		minio: minioclient,
	}
}

func (repo *repository) UploadFile(ctx *fasthttp.RequestCtx, objectName string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	bucket := os.Getenv("S3_BUCKET")

	_, err = repo.minio.PutObject(ctx, bucket, objectName, src, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return err
	}
	return nil
}
