package storage

import (
	"gobackend/pkg/entities"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type Repository interface {
	UploadFile(ctx *fasthttp.RequestCtx, objectname string, file *multipart.FileHeader) error
	AddFileRecord(filename string, path string, size int64, contentType string, Creator uuid.UUID) error
}

type repository struct {
	minio *minio.Client
	db    *gorm.DB
}

func NewRepo(minioclient *minio.Client, db *gorm.DB) Repository {
	return &repository{
		minio: minioclient,
		db:    db,
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
func (repo *repository) AddFileRecord(filename string, path string, size int64, contentType string, Creator uuid.UUID) error {
	return repo.db.Create(&entities.Files{Name: filename, Path: path, Size: size, Type: contentType, Created_by: Creator}).Error
}
