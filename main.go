package main

import (
	"context"
	"net/http"

	"io/ioutil"
	"log"


	// "github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	minio "netlify/minio"
	minioFunc "netlify/minioFunctions"
)
func main() {
	minio.RunMinio()

	e := echo.New()

  	// Middleware
	e.Use(middleware.Logger())
   	e.Use(middleware.Recover())

  	// Routes
  	e.GET("/*", handle)

  	// Start server
  	e.Logger.Fatal(e.Start(":8000"))	
}
func handle(c echo.Context) error {
	res := minioFunc.GetObject(context.Background(), minio.MinioClient, c)
	ans, err := ioutil.ReadAll(res)
	if err != nil {
		log.Println(err)
		return c.Blob(404,"application/octet-stream", ans)
	}
	return c.Blob(http.StatusOK,"application/octet-stream", ans)
}