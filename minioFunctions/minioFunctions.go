package minioFunctions

import (
	"context"
	"fmt"
	"log"
	"strings"

	// "strings"

	// "io"
	// "os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"

//	"net/http"
)

const index = "index.html"

func ListObjects(ctx context.Context, minioClient *minio.Client, backetName string) {
	objectCh := minioClient.ListObjects(context.Background(), backetName, minio.ListObjectsOptions{WithVersions: false, Recursive: true})

	for obj := range(objectCh) {
		if obj.Err != nil {
			log.Println(obj.Err)
		}
		fmt.Println(obj.Key, obj.IsDeleteMarker)
	}
}

func parsePath(path string) string  {
	if path == "" {
		return "index.html"
	}
	parsedPath := strings.Split(filepath.Clean(path), "/")
	correctPath := ""
	for _, v := range(parsedPath) {
        if len(v) > 0 {
            correctPath += v + "/"
        }
    }
	if strings.Contains(correctPath, ".") {
		correctPath = correctPath[:len(correctPath) - 1]
	} else {
		correctPath += "index.html"
	}

	return correctPath
}

func GetObject(ctx context.Context, minioClient *minio.Client, c echo.Context) *minio.Object{
	opts := minio.GetObjectOptions{}
	backetName := strings.Split(c.Request().Host, ":")[0]
	filePath := parsePath(c.Request().URL.Path)
	file, err := minioClient.GetObject(context.Background(), backetName, filePath, opts)
	if err != nil {
		log.Print(err)
	}
	return file
}
