package storage

import (
	"context"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"reflect"
	"testing"
)

func TestMemory_Start(t *testing.T) {
	s := Memory{}
	if err := s.Start(); err != nil {
		t.Fatalf("unexpected error while inititlizing map: %v", err)
	}

	if s.Map == nil {
		t.Fatal("map wasn't initialized correctly")
	}
}

func TestMemory_Create(t *testing.T) {
	s := Memory{MaxSize: 2}
	if err := s.Start(); err != nil {
		t.Fatalf("unexpected error while inititlizing map: %v", err)
	}

	c := v1.Client{Id: "1", Name: "John Doe", Email: "john@email.com", Mobile: "123456"}

	err := s.Create(context.TODO(), &v1.CreateRequest{Client: &c})
	if err != nil {
		t.Fatalf("unexpected error while creating record: %v", err)
	}

	if !reflect.DeepEqual(s.Map[c.Id], c) {
		t.Fatalf("invalid value in the map, expected: %v, got: %v", c, s.Map[c.Id])
	}
}

func TestMemory_Close(t *testing.T) {
	s := Memory{MaxSize: 2}
	if err := s.Start(); err != nil {
		t.Fatalf("unexpected error while inititlizing map: %v", err)
	}

	if err := s.Create(context.TODO(), &v1.CreateRequest{Client: &v1.Client{Id: "1", Name: "John Doe", Email: "john@email.com", Mobile: "123456"}}); err != nil {
		t.Fatalf("unexpected error while creating record: %v", err)
	}

	if err := s.Close(); err != nil {
		t.Fatalf("unexpected error while closing map: %v", err)
	}

	if s.Map != nil {
		t.Fatal("map was not closed properly")
	}
}

func TestMemory_MemoryLimit(t *testing.T) {
	limit := 1
	s := Memory{MaxSize: limit}
	if err := s.Start(); err != nil {
		t.Fatalf("unexpected error while inititlizing map: %v", err)
	}

	if err := s.Create(context.TODO(), &v1.CreateRequest{Client: &v1.Client{Id: "1", Name: "John Doe", Email: "john@email.com", Mobile: "123456"}}); err != nil {
		t.Fatalf("unexpected error while creating record: %v", err)
	}

	if err := s.Create(context.TODO(), &v1.CreateRequest{Client: &v1.Client{Id: "2", Name: "Jane Doe", Email: "jane@email.com", Mobile: "123456"}}); err == nil {
		t.Fatalf("memory limit error expected, but not returned, memory limit: %v", limit)
	}
}
