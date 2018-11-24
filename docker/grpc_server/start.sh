#!/usr/bin/env bash

#wait until mysql is running

while ! mysqladmin ping -h"$DB_HOST" --silent; do
    echo "waiting for MySQL to start"
    sleep 1
done

go install /go/src/github.com/viktorminko/task/pkg/protocol/grpc
go run /go/src/github.com/viktorminko/task/cmd/server/main.go
