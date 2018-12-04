package main

import (
	"log"
	"os"

	"github.com/viktorminko/microservice-task/pkg/cmd/server"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		log.Printf("Error while running server: %v", err)
		os.Exit(1)
	}
}
