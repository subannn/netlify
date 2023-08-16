package server

import (
	"context"
	//"log"

	"mime"
	"strings"

	// "github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	"github.com/minio/minio-go/v7"

	Minio "netlify/minio"
	Logic "netlify/logic"
)
func RunServer(minioClient *minio.Client)  {

	e := echo.New()

  	// Middleware
	e.Use(middleware.Logger())
   	e.Use(middleware.Recover())

  	// Routes
  	e.GET("/*", Handle)

  	// Start server
  	e.Logger.Fatal(e.Start(":8000"))	
}
func Handle(c echo.Context) error {
	p := Logic.ParsePath(c.Request().URL.Path)
	res := Minio.GetObject(context.Background(), Minio.RunMinio(), c, p)


	parts := strings.Split(p, ".")
	fileType := mime.TypeByExtension("." + parts[len(parts) - 1])

	if(fileType == "") {
		fileType = "application/octet"
	}
	return c.Stream(200, fileType , res)
}