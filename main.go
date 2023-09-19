package main

import (
	"os"

	"github.com/Inalegwu/Clipsync/node"
)

func main() {
	nodeType := os.Getenv("TYPE")

	node := node.New()
	if nodeType == "server" {
		node.AsServer()
	} else {
		node.AsClient()
	}
}
