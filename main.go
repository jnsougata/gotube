package main

import (
	"fmt"
	"github.com/gotube/channel"
)

func main() {
	ch := channel.New("https://www.youtube.com/c/SpaceVideosHD")
	for _, v := range ch.Playlists() {
		fmt.Println("....", v)
	}
}
