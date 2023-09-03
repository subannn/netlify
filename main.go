package main

import (
	Minio "netlify/minio"
	Server "netlify/server"
)
func main() {
	Minio.RunMinio()
	Server.RunServer()
}