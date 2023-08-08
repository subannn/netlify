package main

import (
	"context"
	// "fmt"
	"net/http"

	// "fmt"

	"io/ioutil"
	"log"

	// "io"
	// "os"
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
  	e.Logger.Fatal(e.Start(":1323"))
	
}



func handle(c echo.Context) error {
	res := minioFunc.GetObject(context.Background(), minio.MinioClient, c.Request().URL.Path)
	ans, err := ioutil.ReadAll(res)
	if err != nil {
		log.Println(err)
	}
	return c.HTML(http.StatusOK, string(ans))
}