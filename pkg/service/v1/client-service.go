package v1

import (
	"context"
	"database/sql"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type ClientServiceServer struct {
	db *sql.DB
}

func NewClientServiceServer(db *sql.DB) v1.ClientServiceServer {
	return &ClientServiceServer{db: db}
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

func (s *ClientServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

func (s *ClientServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	res, err := c.ExecContext(ctx,
		"INSERT INTO Client(`id`, `Name`, `Email`, `Mobile`) VALUES(?, ?, ?, ?) "+
			"ON DUPLICATE KEY UPDATE `Name`=?, `Email`=?, `Mobile`=?",
		req.Client.Id,
		req.Client.Name,
		req.Client.Email,
		req.Client.Mobile,
		req.Client.Name,
		req.Client.Email,
		req.Client.Mobile)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Client-> "+err.Error())
	}

	// get ID of creates Client
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Client-> "+err.Error())
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}
