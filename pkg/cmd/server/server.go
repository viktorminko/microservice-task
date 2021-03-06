package cmd

import (
	"context"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/viktorminko/microservice-task/pkg/protocol/grpc"
	"github.com/viktorminko/microservice-task/pkg/service/v1"
	"github.com/viktorminko/microservice-task/pkg/storage"
)

const maxMapSize = 10

//InitStorage initializes storage backend
func InitStorage() storage.Storager {
	switch os.Getenv("STORE_TYPE") {
	case "memory":
		return &storage.Memory{MaxSize: maxMapSize}
	}

	return &storage.SQL{}
}

//RunServer runs grpc server
func RunServer() error {

	s := InitStorage()

	err := s.Start()
	if err != nil {
		return err
	}

	defer func() {
		if err := s.Close(); err != nil {
			log.Printf("error while closing storage backend: %v", err)
		}
	}()

	ctx := context.Background()

	v1API := v1.NewClientServiceServer(s)

	p := os.Getenv("PORT")
	return grpc.RunServer(ctx, v1API, p)
}
