package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func RunMinio() {
	endpoint := "192.168.0.106:9000" 
	accessKeyID := "TDIF7k0TmCXXClbV1fil"
	secretAccessKey := "XSFbrmbtyKKGOI5QzlITE9tGpQ9Peagu4ON5FkA4"
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