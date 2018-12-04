package parser

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"io"
	"reflect"
	"strings"
	"testing"
)

type TestData struct {
	Name            string
	data            string
	exp             *v1.Client
	isErrorExpected bool
}

func Parse(newParser func(r io.Reader) Parser, tests []TestData, t *testing.T) {
	for _, tcase := range tests {
		t.Run(tcase.Name, func(t *testing.T) {
			p := newParser(strings.NewReader(tcase.data))
			res, err := p.Parse()

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

func ParseMultipleLines(newParser func(r io.Reader) Parser, s string, exp []*v1.Client, t *testing.T) {

	p := newParser(strings.NewReader(s))

	res, err := p.Parse()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(res, exp[0]) {
		t.Fatalf("unexpected result: expected %v, got %v", exp[0], res)
	}

	res, err = p.Parse()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(res, exp[1]) {
		t.Fatalf("unexpected result: expected %v, got %v", exp[1], res)
	}

	_, err = p.Parse()

	if err == nil {
		t.Fatal("error expected, but not returned")
	}

	if err != io.EOF {
		t.Fatalf("EOF error expected, but got: %v", err)
	}
}
