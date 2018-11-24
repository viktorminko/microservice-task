package main

import (
	"encoding/csv"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"reflect"
	"strings"
	"testing"
)

func TestParseClient(t *testing.T) {

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
			&v1.Client{Name: "Kieran", Email: "Cras@magna.ca", Mobile: "+440128568417"},
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
			&v1.Client{Name: "Kieran", Email: "Cras@magna.ca", Mobile: "+440128568417"},
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
			reader := csv.NewReader(strings.NewReader(tcase.data))
			reader.TrimLeadingSpace = true
			res, err := ParseClient(*reader)

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
