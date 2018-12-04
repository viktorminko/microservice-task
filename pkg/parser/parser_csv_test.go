package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestCSV_Parse(t *testing.T) {

	type TestData struct {
		Name            string
		data            string
		exp             *v1.Client
		isErrorExpected bool
	}

	var tests = []TestData{
		{
			"valid client",
			`39,Kieran,Cras@magna.ca,(01285) 68417`,
			&v1.Client{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "+440128568417"},
			false,
		},
		{
			"invalid_fields_number",
			`39,Kieran, UK, Cras@magna.ca,(01285) 68417`,
			nil,
			true,
		},
		{
			"phone number formatted correctly",
			`39,Kieran, Cras@magna.ca,+s(01285) 68417ml`,
			&v1.Client{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "+440128568417"},
			false,
		},
		{
			"first field should be an integer",
			`id,name,email,mobile_number`,
			nil,
			true,
		},
	}

	for _, tcase := range tests {
		t.Run(tcase.Name, func(t *testing.T) {
			p := CSV{}
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

func TestCSV_ParseMultipleLines(t *testing.T) {

	r := strings.NewReader(`
		39,Kieran,Cras@magna.ca,(01285) 68417
    	2,John,johndoe@gmail.com,123456`)

	exp := []*v1.Client{
		{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "+440128568417"},
		{Id: "2", Name: "John", Email: "johndoe@gmail.com", Mobile: "+44123456"},
	}

	p := CSV{}

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
