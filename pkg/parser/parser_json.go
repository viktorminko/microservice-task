package parser

import (
	"bufio"
	"encoding/json"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
)

//JSON parses json strings from reader line by line
//each line should be a valid json string
type JSON struct {
	Scanner *bufio.Scanner
}

//NewJSON creates new JSON parser that reads from r
func NewJSON(r io.Reader) *JSON {
	return &JSON{bufio.NewScanner(r)}
}

//Parse parses single json entry representing one client
func (p *JSON) Parse() (*v1.Client, error) {
	if !p.Scanner.Scan() {
		if p.Scanner.Err() != nil {
			return nil, p.Scanner.Err()
		}
		return nil, io.EOF
	}

	res := &v1.Client{}
	err := json.Unmarshal([]byte(p.Scanner.Text()), res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
