package main

import (
	"log"

	"github.com/ajalck/spine/pkg/config"
	"github.com/ajalck/spine/pkg/di"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load Config: %v", err)
	}

	server, err := di.InitializeAPI(c)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
