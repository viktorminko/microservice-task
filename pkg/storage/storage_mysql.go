package storage

import (
	"database/sql"
	"fmt"
	"github.com/viktorminko/microservice-task/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

//SQL implements MySQL storage
type SQL struct {
	Db *sql.DB
}

//Config handles server configuration
type config struct {
	Port string

	DBHost     string
	DBUser     string
	DBPassword string
	DBDatabase string
}

const dbMaxOpenConnections = 100

//Start opens DB connection
func (s *SQL) Start() error {
	// get configuration
	var cfg config
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBDatabase = os.Getenv("DB_SCHEMA")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBDatabase,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(dbMaxOpenConnections)

	s.Db = db

	return nil
}

func (s *SQL) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.Db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

//Create creates record in MySQL DB
func (s *SQL) Create(ctx context.Context, req *v1.CreateRequest) error {
	// get SQL connection from pool
	c, err := s.connect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := c.Close(); err != nil {
			log.Printf("error while closing mysql connection: %v", err)
		}
	}()

	_, err = c.ExecContext(ctx,
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
		return status.Error(codes.Unknown, "failed to insert into Client-> "+err.Error())
	}

	return nil
}

//Close closes MySQL DB
func (s *SQL) Close() error {
	log.Printf("Closing db")
	return s.Db.Close()
}
