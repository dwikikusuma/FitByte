package repositories

import "github.com/minio/minio-go/v7"

type MinioRepository interface {
}

type minioRepository struct {
	client *minio.Client
}

func NewMinoRepository(storageClient *minio.Client) MinioRepository {
	return &minioRepository{
		client: storageClient,
	}
}
