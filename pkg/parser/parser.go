package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
)

//to add new parser create struct which implements this interface
type Parser interface {
	Parse(r io.Reader) (*v1.Client, error)
}
