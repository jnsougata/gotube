package main

import (
	"fmt"
	"github.com/gotube/channel"
)

func main() {
	ch := channel.New("https://www.youtube.com/c/livenowfox")
	fmt.Println(ch.Info())
}
