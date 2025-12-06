package main

import (
	"context"
	"fmt"

	transmission "github.com/j-dumbell/go-qbittorrent"
)

func main() {
	ctx := context.Background()

	client, err := transmission.New(transmission.ClientParams{
		Host:     "http://localhost:9091",
		User:     "admin",
		Password: "password",
	})
	if err != nil {
		panic(err)
	}

	session, err := client.SessionGet(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(session)
}
