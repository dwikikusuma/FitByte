package service

import "FitByte/internal/repositories"

type FileService interface {
}

type fileService struct {
	fileRepo  repositories.FileRepository
	minioRepo repositories.MinioRepository
}

func NewFileService(fileRepo repositories.FileRepository, storageRepo repositories.MinioRepository) FileService {
	return &fileService{
		fileRepo:  fileRepo,
		minioRepo: storageRepo,
	}
}
