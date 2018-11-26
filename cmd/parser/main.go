package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/viktorminko/microservice-task/pkg/api/v1"
)

const (
	apiVersion       = "v1"
	executionTimeout = 100 * time.Second
)

type Client struct {
	Name         string
	Email        string
	MobileNumber string
}

func ParseClient(reader csv.Reader) (*v1.Client, error) {
	line, err := reader.Read()

	if err != nil {
		return nil, err
	}

	if len(line) != 4 {
		return nil, fmt.Errorf(
			"unexpected number of properties in client line: expected %v, got %v",
			4, len(line),
		)
	}

	//First field should be an integer
	_, err = strconv.Atoi(line[0])
	if err != nil {
		return nil, err
	}

	//Leave only numbers in the phone
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return nil, err
	}

	mobile := "+44" + reg.ReplaceAllString(
		strings.Replace(line[3], " ", "", -1),
		"",
	)

	return &v1.Client{Id: line[0], Name: line[1], Email: line[2], Mobile: mobile}, nil
}

func runRequest(client *v1.Client, c v1.ClientServiceClient, ctx context.Context) (*v1.CreateResponse, error) {
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

func main() {

	c, conn, err := initClientService(os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT"))
	defer conn.Close()

	if err != nil {
		log.Fatalf("unable to init client service: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), executionTimeout)
	defer cancel()

	csvFile, err := os.Open(os.Getenv("DATA_FILE"))

	if err != nil {
		log.Fatalf("unable to open CSV file: %v", err)
	}

	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true

	var wg sync.WaitGroup

	for {

		client, err := ParseClient(*reader)
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

			res, err := runRequest(client, c, ctx)
			if err != nil {
				log.Printf("error while sending request to client service: %v", err)
				return
			}

			if 0 == res.Id {
				log.Println("record updated\n")
			} else {
				log.Printf("new record added with id: %v \n\n", res.Id)
			}

		}()

		wg.Wait()
	}

}
