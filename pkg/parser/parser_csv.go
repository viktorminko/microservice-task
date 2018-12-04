package parser

import (
	"encoding/csv"
	"fmt"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
	"regexp"
	"strconv"
	"strings"
)

//CSV parses csv data from Reader line by line
type CSV struct {
	Reader *csv.Reader
}

//Parse executes paring of one client
func (p *CSV) Parse(r io.Reader) (*v1.Client, error) {
	if p.Reader == nil {
		p.Reader = csv.NewReader(r)
		p.Reader.TrimLeadingSpace = true
	}

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

	//Leave only numbers in the phone
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return nil, err
	}

	mobile := "+44" + reg.ReplaceAllString(
		strings.Replace(res[3], " ", "", -1),
		"",
	)

	return &v1.Client{Id: res[0], Name: res[1], Email: res[2], Mobile: mobile}, nil
}
