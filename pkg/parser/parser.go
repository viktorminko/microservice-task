package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
)

//Parser interface for parsers
type Parser interface {
	Parse(r io.Reader) (*v1.Client, error)
}
