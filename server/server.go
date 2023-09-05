package server

import (
	"mime"
	"strings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subannn/netlify/logic"
	"github.com/subannn/netlify/minio"
)

func RunServer() {
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
	p := logic.ParsePath(c.Request().URL.Path)
	res := minio.GetObject(c, p)
	parts := strings.Split(p, ".")
	fileType := mime.TypeByExtension("." + parts[len(parts)-1])
	if fileType == "" {
		fileType = "application/octet"
	}
	return c.Stream(200, fileType, res)
}
