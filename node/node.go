package node

import (
	"fmt"

	"github.com/Inalegwu/Clipsync/data"
	"github.com/Inalegwu/Clipsync/message"
	"github.com/google/uuid"
)

type Node struct {
	Id          uuid.UUID
	connections []string
	data        data.Data
}

func New() Node {
	id := uuid.New()

	return Node{
		Id:          id,
		connections: make([]string, 5),
	}
}

func (n *Node) HandleStream() {
}

func (n *Node) handleMessage(message message.Message) {
	messageAsJson := message.ToJson()

	fmt.Println(messageAsJson)
}

func (n *Node) sendMessage(message_type int, receiver_id uuid.UUID) {
	message := message.New(n.Id, receiver_id, message_type)

	message.ToJson()
}
