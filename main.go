package main

import (
	"context"
	//"log"

	"mime"
	"strings"

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
	p := minioFunc.ParsePath(c.Request().URL.Path)
	res := minioFunc.GetObject(context.Background(), minio.MinioClient, c, p)


	parts := strings.Split(p, ".")
	fileType := mime.TypeByExtension("." + parts[len(parts) - 1])

	if(fileType == "") {
		fileType = "application/octet"
	}
	return c.Stream(200, fileType , res)
}