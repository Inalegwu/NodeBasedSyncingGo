package main

import (
	"log"
	"os"

	"github.com/Inalegwu/Clipsync/node"
)

func main() {
	nodeType := os.Getenv("TYPE")

	node := node.New()
	if nodeType == "server" {
		node.AsServer()
	} else if nodeType == "client" {
		node.AsClient()
	} else {
		log.Fatal("provide a node type to start the server TYPE='server' or 'client'")
	}
}
