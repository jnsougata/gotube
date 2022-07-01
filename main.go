package main

import (
	"fmt"
	"github.com/gotube/search"
)

func main() {
	ch := search.New("pwedipie").Channel()
	fmt.Println(ch.Info())
}
