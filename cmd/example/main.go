package main

import (
	"context"

	"github.com/j-dumbell/go-qbittorrent"
)

func main() {
	ctx := context.Background()

	client, err := qbittorrent.New(qbittorrent.ClientParams{
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
