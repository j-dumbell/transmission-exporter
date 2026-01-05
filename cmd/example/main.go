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

	// torrents, err := client.TorrentGet(ctx, transmission.AllTorrents)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// var torrentID int64
	// for _, torrent := range torrents.Torrents {
	// 	if torrent.Name == "ubuntu-25.10-desktop-amd64.iso" {
	// 		torrentID = torrent.ID
	// 	}
	// }
	// fmt.Println("torrentID ", torrentID)
	//
	// if err := client.QueueMoveTop(ctx, transmission.NewTorrentIDs(torrentID)); err != nil {
	// 	panic(err)
	// }

	groups, err := client.GroupGet(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", groups)

	// err = client.GroupSet(ctx, transmission.GroupSetArgs{
	// 	Name:                  "blahblah",
	// 	HonorsSessionLimits:   toPtr(true),
	// 	SpeedLimitDown:        toPtr(10),
	// 	SpeedLimitDownEnabled: toPtr(false),
	// 	SpeedLimitUp:          toPtr(20),
	// 	SpeedLimitUpEnabled:   toPtr(false),
	// })
	// if err != nil {
	// 	panic(err)
	// }
}

func toPtr[T any](t T) *T {
	return &t
}
