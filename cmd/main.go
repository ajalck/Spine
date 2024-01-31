package main

import (
	"fmt"
	"log"

	"github.com/ajalck/spine/pkg/di"
)

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()
	server := di.InitializeAPI()
	fmt.Println(server)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
