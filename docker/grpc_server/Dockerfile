FROM golang:latest

#We need this to check MySQL server is running
RUN apt-get update
RUN apt-get install -y mysql-client

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/go-sql-driver/mysql