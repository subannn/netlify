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


func RunMinio() *minio.Client{
	endpoint := os.Getenv("ENDPOINT")            // "192.168.0.106:9000" 
	accessKeyID := os.Getenv("ACCESS_KEY_ID")  		  // "TDIF7k0TmCXXClbV1fil"
	secretAccessKey := os.Getenv("SECRET_ACCESS_KEY")  		  // "XSFbrmbtyKKGOI5QzlITE9tGpQ9Peagu4ON5FkA4"
	useSSL := false
	// initialiRunMiniozation of Minio Client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}

func GetObject(ctx context.Context, minioClient *minio.Client, c echo.Context, filePath string) *minio.Object{
	opts := minio.GetObjectOptions{}
	backetName := strings.Split(c.Request().Host, ":")[0]

	file, err := minioClient.GetObject(context.Background(), backetName, filePath, opts)
	if err != nil {
		log.Print(err)
	}
	return file
}
