package minioFunctions

import (
	"context"
	"fmt"
	"log"
	"strings"

	// "io"
	// "os"

	"github.com/minio/minio-go/v7"
	//"github.com/minio/minio-go/v7/pkg/credentials"
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

func parsePath(path string) (string, string)  {
	s := strings.Split(path, "/")
	bucketName := s[1]
	pathName := ""
	for i := 2;i < len(s);i++ {
		pathName += s[i] + "/"
	}
	return bucketName, pathName + "index.html"
}

func GetObject(ctx context.Context, minioClient *minio.Client, path string) *minio.Object{
	opts := minio.GetObjectOptions{}
	backetName, filePath := parsePath(path)
	file, err := minioClient.GetObject(ctx, backetName, filePath, opts)
	if err != nil {
		log.Print(err)
	}
	return file
}
