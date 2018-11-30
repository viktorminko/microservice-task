package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/viktorminko/microservice-task/pkg/protocol/grpc"
	"github.com/viktorminko/microservice-task/pkg/service/v1"
)

//Config handles server configuration
type Config struct {
	Port string

	DBHost     string
	DBUser     string
	DBPassword string
	DBDatabase string
}

const dbMaxOpenConnections = 100

//RunServer runs grpc server
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	cfg.Port = os.Getenv("PORT")
	cfg.DBHost = os.Getenv("DB_HOST")
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBDatabase = os.Getenv("DB_SCHEMA")

	if len(cfg.Port) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.Port)
	}

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
	defer db.Close()

	db.SetMaxOpenConns(dbMaxOpenConnections)

	v1API := v1.NewClientServiceServer(db)

	return grpc.RunServer(ctx, v1API, cfg.Port)
}
