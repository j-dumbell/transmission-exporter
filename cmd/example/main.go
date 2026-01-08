package main

import (
	"context"
	"fmt"

	"github.com/j-dumbell/go-qbittorrent/pkg/transmission"
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

	fmt.Printf("%+v\n", session)
}
