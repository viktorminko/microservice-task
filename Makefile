 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_SERVER=server
BINARY_NAME_PARSER=parser

build: build_server build_parser

run: run_server run_parser

all: test build

run_all: run_server run_parser

test:
		$(GOTEST) -v ./...

build_server:
		$(GOBUILD) -o $(BINARY_NAME_SERVER) -v ./cmd/server/...

run_server:
		$(GOBUILD) -o $(BINARY_NAME_SERVER) -v ./cmd/server/...
		PORT=5000 DBHOST=mysql DB_USER=user DB_PASSWORD=password DB_SCHEMA=schema ./$(BINARY_NAME_SERVER)&

build_parser:
		$(GOBUILD) -o $(BINARY_NAME_PARSER) -v ./cmd/parser/...

run_parser:
		$(GOBUILD) -o $(BINARY_NAME_PARSER) -v ./cmd/parser/...
		GRPC_HOST=0.0.0.0 GRPC_PORT=5000 DATA_FILE=./data.csv ./$(BINARY_NAME_PARSER)

clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
