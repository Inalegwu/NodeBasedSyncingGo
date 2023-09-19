package main

import (
	"fmt"
	"os"

	"github.com/Inalegwu/Clipsync/node"
)

func main() {
	nodeType := os.Getenv("TYPE")

	node := node.New()
	if nodeType == "server" {
		fmt.Println("Starting Node as Server")

		fmt.Printf("node with id : %v created \n", node.Id)
		node.AsServer()
	} else {
		fmt.Println("Starting node as client")

		fmt.Printf("node with id : %v created \n", node.Id)
		node.AsClient()
	}
}
