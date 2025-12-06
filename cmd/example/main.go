package main

import (
	"context"

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

	if err := client.SessionGet(ctx); err != nil {
		panic(err)
	}
}
