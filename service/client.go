package service

import (
	"context"
	"fmt"
	"time"

	"github.com/peng225/wolper/pb"
	"google.golang.org/grpc"
)

func ClientQuery(addrAndPort, key, include, exclude string, uniq, entropySort bool) []string {
	conn, err := grpc.Dial(
		addrAndPort,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		fmt.Println("Connection failed.")
		return nil
	}
	defer conn.Close()
	fmt.Println("Connection succeeded.")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	client := pb.NewWolperServiceClient(conn)
	searchRequest := pb.SearchRequest{
		Key:         key,
		Include:     include,
		Exclude:     exclude,
		Uniq:        uniq,
		EntropySort: entropySort,
	}

	result, err := client.Query(ctx, &searchRequest)
	if err != nil {
		fmt.Println("Request failed.")
		return nil
	}

	return result.Words
}
