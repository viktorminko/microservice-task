package main

import (
	"bufio"
	"context"
	"github.com/viktorminko/microservice-task/pkg/parser"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/viktorminko/microservice-task/pkg/api/v1"
)

const (
	apiVersion       = "v1"
	executionTimeout = 100 * time.Second
)

func runRequest(ctx context.Context, client *v1.Client, c v1.ClientServiceClient) (*v1.CreateResponse, error) {
	req := v1.CreateRequest{
		Api:    apiVersion,
		Client: client,
	}

	res, err := c.Create(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func initClientService(host string, port string) (v1.ClientServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return v1.NewClientServiceClient(conn), conn, nil
}

//init reader for data source
//can be changed if we need to reed from something else then a file
func initSource() (io.Reader, error) {
	f, err := os.Open(os.Getenv("DATA_FILE"))

	if err != nil {
		return nil, err
	}

	return bufio.NewReader(f), nil
}

//init
func initParser() parser.Parser {
	switch os.Getenv("DATA_TYPE") {
	case "csv":
		return &parser.CSV{}
	case "json":
		return &parser.JSON{}
	}

	return &parser.CSV{}
}

func main() {

	c, conn, err := initClientService(os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT"))
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("error while closing client service connection: %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("unable to init client service: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), executionTimeout)
	defer cancel()

	src, err := initSource()
	if err != nil {
		log.Fatalf("unable to init data source: %v", err)
	}

	p := initParser()

	var wg sync.WaitGroup
	for {
		client, err := p.Parse(src)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("unable to parse client record: %v", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			log.Printf("sending request to client service, record: %v", client)

			_, err := runRequest(ctx, client, c)
			if err != nil {
				log.Printf("error while sending request to client service: %v", err)
				return
			}

			log.Println("record was updated")

		}()
	}

	wg.Wait()

}
