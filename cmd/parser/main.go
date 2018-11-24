package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/viktorminko/task/pkg/api/v1"
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

func main() {

	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewClientServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), executionTimeout)
	defer cancel()

	csvFile, err := os.Open(os.Getenv("DATA_FILE"))

	if err != nil {
		log.Fatal("unable to open csv file", err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	//Skip first line
	_, err = reader.Read()
	if err != nil {
		log.Fatal("invalid first line", err)
	}

	var wg sync.WaitGroup

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		//Leave only numbers in the phone
		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			log.Fatal(err)
		}

		mobile := "+44" + reg.ReplaceAllString(
			strings.Replace(line[3], " ", "", -1),
			"",
		)

		wg.Add(1)
		go func() {
			defer wg.Done()
			req := v1.CreateRequest{
				Api: apiVersion,
				Client: &v1.Client{
					Name:   line[1],
					Email:  line[2],
					Mobile: mobile,
				},
			}

			log.Printf("adding new record: %v", line)

			res1, err := c.Create(ctx, &req)
			if err != nil {
				log.Fatalf("create failed: %v", err)
			}
			log.Printf("create result: <%+v>\n\n", res1)
		}()

		wg.Wait()
	}

}
