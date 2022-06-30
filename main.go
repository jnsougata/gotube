package main

import (
	"fmt"
	"github.com/gotube/search"
)

func main() {
	ch := search.New("YouTube Shorts").Channel()
	fmt.Println(ch.Info())
}
