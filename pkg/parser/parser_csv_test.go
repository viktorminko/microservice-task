package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
	"testing"
)

func TestCSV_Parse(t *testing.T) {

	Parse(
		func(r io.Reader) Parser { return NewCSV(r) },
		[]TestData{
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
		},
		t,
	)
}

func TestCSV_ParseMultipleLines(t *testing.T) {

	ParseMultipleLines(
		func(r io.Reader) Parser { return NewCSV(r) },
		`
		39,Kieran,Cras@magna.ca,(01285) 68417
    	2,John,johndoe@gmail.com,123456`,
		[]*v1.Client{
			{Id: "39", Name: "Kieran", Email: "Cras@magna.ca", Mobile: "+440128568417"},
			{Id: "2", Name: "John", Email: "johndoe@gmail.com", Mobile: "+44123456"},
		},
		t,
	)
}
