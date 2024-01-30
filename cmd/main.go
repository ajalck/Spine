package main

import (
	"log"

	"github.com/ajalck/spine/pkg/di"
)

func main() {
	server := di.InitializeAPI()
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
