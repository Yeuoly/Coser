package db

import (
	"context"
	"fmt"

	"github.com/Yeuoly/billboards/internal/static"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioEndpoint  string
	minioAccessKey string
	minioSecretKey string
	minioBucket    string
	minioUseSSL    = false

	client *minio.Client
)

func InitMinio() error {
	var err error

	config := static.GetBillboardsGlobalConfigurations()
	minioEndpoint = fmt.Sprintf("%s:%d", config.Minio.Host, config.Minio.Port)
	minioAccessKey = config.Minio.AK
	minioSecretKey = config.Minio.SK
	minioBucket = config.Minio.Bucket

	client, err = minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: minioUseSSL,
	})

	if err != nil {
		return err
	}

	// check if bucket exists
	exists, err := client.BucketExists(context.Background(), minioBucket)
	if err != nil {
		return err
	}

	if !exists {
		err = client.MakeBucket(context.Background(), minioBucket, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
