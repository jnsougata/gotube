package main

import (
	"fmt"
	"github.com/gotube/channel"
	// "github.com/gotube/search"
)

func main() {
	ch := channel.New("https://www.youtube.com/c/GolangDojo")
	fmt.Println(ch.LatestUploaded())
}
