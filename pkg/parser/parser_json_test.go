package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestJSON_Parse(t *testing.T) {

	type TestData struct {
		Name            string
		data            string
		exp             *v1.Client
		isErrorExpected bool
	}

	var tests = []TestData{
		{
			"valid client",
			`{"Id": "39","Name": "Kieran","Email":"Cras@magna.ca","Mobile": "(01285) 68417"}`,
			&v1.Client{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "(01285) 68417"},
			false,
		},
	}

	for _, tcase := range tests {
		t.Run(tcase.Name, func(t *testing.T) {
			p := JSON{}
			res, err := p.Parse(strings.NewReader(tcase.data))

			if tcase.isErrorExpected && err == nil {
				t.Fatal("error expected, but not returned")
			}

			if !tcase.isErrorExpected && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(res, tcase.exp) {
				t.Fatalf("unexpected result: expected %v, got %v", tcase.exp, res)
			}
		})

	}
}

func TestJson_ParseMultipleLines(t *testing.T) {

	r := strings.NewReader(`{"Id": "39","Name": "Kieran","Email":"Cras@magna.ca",    "Mobile": "(01285) 68417"}
    	{"Id": "2", "Name": "John",  "Email":"johndoe@gmail.com","Mobile": "12345"}`)

	exp := []*v1.Client{
		{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "(01285) 68417"},
		{Id: "2", Name: "John", Email: "johndoe@gmail.com", Mobile: "12345"},
	}

	p := JSON{}

	res, err := p.Parse(r)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(res, exp[0]) {
		t.Fatalf("unexpected result: expected %v, got %v", exp[0], res)
	}

	res, err = p.Parse(r)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(res, exp[1]) {
		t.Fatalf("unexpected result: expected %v, got %v", exp[1], res)
	}

	_, err = p.Parse(r)

	if err == nil {
		t.Fatal("error expected, but not returned")
	}

	if err != io.EOF {
		t.Fatalf("EOF error expected, but got: %v", err)
	}
}
