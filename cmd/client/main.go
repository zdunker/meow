package main

import (
	"fmt"

	"github.com/zdunker/meow/client"
)

func main() {
	client.Table.Init()
	fmt.Println(client.Table)
	fmt.Println(client.Table.Length())
}
