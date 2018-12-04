package storage

import (
	"context"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
)

//Storager interface for different storage backends
type Storager interface {
	Start() error
	Create(ctx context.Context, req *v1.CreateRequest) error
	Close() error
}
