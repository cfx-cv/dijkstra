package main

import (
	"os"

	"../pkg/server"
)

func main() {
	server.Start(os.Getenv("PORT"), os.Getenv("API_KEY"))
}
