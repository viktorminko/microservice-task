FROM golang

#Required to check if grpc_server container is running
RUN apt-get update
RUN apt-get install -y netcat

RUN go get -u github.com/golang/protobuf/proto
RUN go get -u google.golang.org/grpc