package main

import (
	Minio "github.com/subannn/netlify/minio"
	Server "github.com/subannn/netlify/server"
)

func main() {
	Minio.RunMinio()
	Server.RunServer()
}
