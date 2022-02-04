package service

import (
	"fmt"
	"context"
	"time"
    "github.com/peng225/wolper/pb"
    "google.golang.org/grpc"
)

func ClientQuery(addrAndPort, include, exclude, fixed string) {
    fmt.Printf("service.ClientQeury called (addrAndPort = %v).\n", addrAndPort)
    conn, err := grpc.Dial(
        addrAndPort,
        grpc.WithInsecure(),
        grpc.WithBlock(),
    )
    if err != nil {
        fmt.Println("Connection failed.")
        return
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
        Key: fixed,
        Include: include,
        Exclude: exclude,
    }

    result, err := client.Query(ctx, &searchRequest)
    if err != nil {
        fmt.Println("Request failed.")
        return
    }

    words := result.Words
    for _, word := range words {
        fmt.Println(word)
    }
}
