package parser

import (
	"encoding/csv"
	"fmt"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"github.com/viktorminko/microservice-task/pkg/client"
	"io"
	"strconv"
)

//CSV parses csv data from Reader line by line
type CSV struct {
	Reader *csv.Reader
}

//NewCSV creates new CSV parser that reads from r
func NewCSV(r io.Reader) *CSV {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true
	return &CSV{cr}
}

//Parse executes parsing of one client
func (p *CSV) Parse() (*v1.Client, error) {
	res, err := p.Reader.Read()
	if err != nil {
		return nil, err
	}

	if len(res) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of properties in client line: expected %v, got %v",
			4, len(res),
		)
	}

	//First field should be an integer
	_, err = strconv.Atoi(res[0])
	if err != nil {
		return nil, err
	}

	mobile, err := client.NewMobile(res[3]).GetFormattedNumber()
	if err != nil {
		return nil, err
	}

	return &v1.Client{Id: res[0], Name: res[1], Email: res[2], Mobile: mobile}, nil
}
