package storage

import (
	"mime/multipart"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/valyala/fasthttp"
)

type Repository interface {
	UploadFile(ctx *fasthttp.RequestCtx, filename string, file *multipart.FileHeader) error
}

type repository struct {
	minio *minio.Client
}

func NewRepo(minioclient *minio.Client) Repository {
	return &repository{
		minio: minioclient,
	}
}

func (r *repository) UploadFile(ctx *fasthttp.RequestCtx, filename string, file *multipart.FileHeader) error {
	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	r.minio.PutObject(ctx, os.Getenv("S3_BUCKET"), filename, srcFile, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	return nil

}
