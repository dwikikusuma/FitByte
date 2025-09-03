package infrastructure

import (
	"FitByte/configs"
	"FitByte/pkg/log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioStorage(appConfig configs.Config) *minio.Client {
	endpoint := appConfig.Minio.Endpoint
	accessKeyID := appConfig.Minio.AccessKeyID
	secretAccessKey := appConfig.Minio.SecretAccessKey
	useSSL := appConfig.Minio.UseSSL

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("minio init failed")
	}

	log.Logger.Info().Msg("minio init success")
	return minioClient
}
