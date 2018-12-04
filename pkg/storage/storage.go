package storage

import (
	"context"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
)

type Storager interface {
	Start() error
	Create(ctx context.Context, req *v1.CreateRequest) error
	Close() error
}
