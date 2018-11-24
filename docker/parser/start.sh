#!/usr/bin/env bash

#wait until grpc_server is running

while true; do
    nc -zv $GRPC_HOST $GRPC_PORT &> /dev/null ;
    if [ $? -eq 0 ]; then
        break
    fi
    echo "waiting for grpc_server to start"
    sleep 1
done

go run /go/src/github.com/viktorminko/task/cmd/parser
