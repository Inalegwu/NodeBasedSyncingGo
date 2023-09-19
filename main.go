package main

import (
	"fmt"

	"github.com/Inalegwu/Clipsync/node"
)

func main() {
	node := node.New()

	fmt.Printf("node with id : %v created \n", node.Id)
}
