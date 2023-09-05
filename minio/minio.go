package minio

import (
	"context"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func RunMinio() {
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

// Get object from Minio backet
func GetObject(c echo.Context, filePath string) *minio.Object {
	opts := minio.GetObjectOptions{}
	backetName := c.Request().Host
	log.Println(backetName, filePath)
	file, err := minioClient.GetObject(context.Background(), backetName, filePath, opts)
	if err != nil {
		log.Println(err)
	}
	return file
}
