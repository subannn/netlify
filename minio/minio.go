package minio

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func RunMinio() {
	endpoint := os.Getenv("ENDPOINT") 
	accessKeyID := os.Getenv("ACCESS_KEY_ID") 
	secretAccessKey := os.Getenv("SECRET_ACCESS_KEY") 
	useSSL := false
	// initialization of Minio Client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	MinioClient = minioClient
}