package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
	"testing"
)

func TestJSON_Parse(t *testing.T) {
	Parse(
		func(r io.Reader) Parser { return NewJSON(r) },
		[]TestData{
			{
				"valid client",
				`{"Id": "39","Name": "Kieran","Email":"Cras@magna.ca","Mobile": "(01285) 68417"}`,
				&v1.Client{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "(01285) 68417"},
				false,
			},
		},
		t,
	)
}

func TestJson_ParseMultipleLines(t *testing.T) {
	ParseMultipleLines(
		func(r io.Reader) Parser { return NewJSON(r) },
		`{"Id": "39","Name": "Kieran","Email":"Cras@magna.ca",    "Mobile": "(01285) 68417"}
    	{"Id": "2", "Name": "John",  "Email":"johndoe@gmail.com","Mobile": "12345"}`,
		[]*v1.Client{
			{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "(01285) 68417"},
			{Id: "2", Name: "John", Email: "johndoe@gmail.com", Mobile: "12345"},
		},
		t,
	)
}
