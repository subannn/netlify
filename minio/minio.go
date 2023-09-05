package minio

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func RunMinio() {
	/*
		endpoint := "192.168.0.106:9000"
		accessKeyID := "TDIF7k0TmCXXClbV1fil"
		secretAccessKey := "XSFbrmbtyKKGOI5QzlITE9tGpQ9Peagu4ON5FkA4"
	*/
	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("SECRET_ACCESS_KEY")
	log.Println(endpoint, "\n", accessKeyID, "\n", secretAccessKey)
	useSSL := false
	// initialiRunMiniozation of Minio Client
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println(err)
	}
	minioClient = client
}

func GetObject(c echo.Context, filePath string) *minio.Object {
	opts := minio.GetObjectOptions{}
	backetName := strings.Split(c.Request().Host, ":")[0]
	log.Println(backetName, filePath)
	file, err := minioClient.GetObject(context.Background(), backetName, filePath, opts)
	if err != nil {
		log.Println(err)
	}
	return file
}
