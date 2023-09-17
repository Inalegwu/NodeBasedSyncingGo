package node

import (
	"fmt"

	"github.com/Inalegwu/Clipsync/data"
	"github.com/Inalegwu/Clipsync/message"
)

type Node struct {
	id          string
	connections []string
	data        data.Data
}

func New() Node {
	return Node{
		id:          "1",
		connections: make([]string, 5),
	}
}

func (n *Node) handleMessage(message message.Message) {
	messageAsJson := message.ToJson()

	fmt.Println(messageAsJson)
}

func (n *Node) sendMessage(message_type int, receiver_id string) {
	message := message.New(n.id, receiver_id, message_type)

	message.ToJson()
}
