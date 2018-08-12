package main

import (
	"os"

	"github.com/cfx-cv/trail/pkg/server"
)

func main() {
	server.Start(os.Getenv("PORT"), os.Getenv("API_KEY"))
}
