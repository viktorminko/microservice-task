package v1

import (
	"context"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"github.com/viktorminko/microservice-task/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

//ClientServiceServer represents service to handle grpc requests
type ClientServiceServer struct {
	Storage storage.Storager
}

//NewClientServiceServer creates client service
func NewClientServiceServer(s storage.Storager) v1.ClientServiceServer {
	return &ClientServiceServer{s}
}

func (s *ClientServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the parser
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: parser implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

//Create creates new record in storage
func (s *ClientServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	err := s.Storage.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.CreateResponse{
		Api: apiVersion,
	}, nil
}
