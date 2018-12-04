package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
)

//Parser interface for parsers
type Parser interface {
	Parse() (*v1.Client, error)
}
