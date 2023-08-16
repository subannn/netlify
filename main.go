package main

import (
	Minio "netlify/minio"
	Server "netlify/server"
)
func main() {
	minioClient := Minio.RunMinio()
	Server.RunServer(minioClient)
}