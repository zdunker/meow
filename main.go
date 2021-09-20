package main

import (
	"fmt"

	"github.com/zdunker/meow/client"
)

func main() {
	client.MakeTable()
	fmt.Println(client.Table)
	fmt.Println(client.Table.Length())
}
