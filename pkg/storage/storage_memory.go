package storage

import (
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

//Memory stores data in map
type Memory struct {
	Map     map[string]v1.Client
	MaxSize int
	mutex   sync.Mutex
}

//Start initializes map
func (m *Memory) Start() error {
	m.Map = make(map[string]v1.Client)
	m.mutex = sync.Mutex{}

	return nil
}

//Create creates creates record in DB
func (m *Memory) Create(ctx context.Context, req *v1.CreateRequest) error {

	if m.Map == nil {
		return status.Errorf(codes.Unknown, "mao is not initialized")
	}

	if len(m.Map) >= m.MaxSize {
		return status.Errorf(codes.Unknown, "Out of memory, maxMapSize reached: %v", m.MaxSize)
	}

	client := req.Client

	m.mutex.Lock()
	m.Map[client.Id] = *client
	m.mutex.Unlock()

	return nil
}

//Close clears the map
func (m *Memory) Close() error {
	m.Map = nil
	return nil
}
