### Run with docker

I left .env files in git to simplify set up process, normally only .env samples should be in repository

Run containers:

put file data.csv in route folder
run

`
docker-compose up -d
`

### Run compiled binaries

Run server with your DB credentials

`
PORT=5000 DBHOST=mysql DB_USER=user DB_PASSWORD=password DB_SCHEMA=schema go run cmd/server/main.go
`

then run parser

`
 GRPC_HOST=0.0.0.0 GRPC_PORT=5000 DATA_FILE=./data.csv  go run cmd/parser/main.go
`

### Run with Makefile

Test, lint and build

`
make all
`

Run server and parser

`
 make run_all
`

### Directory structure

* /api/proto - proto file for specific api version (current v1)
* /cmd/server/ - server entrypoint (main.go)
* /cmd/parser/ - parser entrypoint (main.go)
* /docker - configurations for docker containers from docker-compose.yml
* /pkg - packages for server and parser
* /pkg/api/v1 - automatically generated code from proto file
* /pkg/cmd/server - grpc server set up configuration
* /pkg/parser - parsers and tests
* /pkg/protocol/grpc - grpc server runner
* /pkg/service/ - service to handle requests from api
* protoc-gen.sh - generate code based on proto file
