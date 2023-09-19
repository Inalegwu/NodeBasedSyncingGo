package node

import (
	"fmt"
	"log"
	"net"

	"github.com/Inalegwu/Clipsync/data"
	"github.com/Inalegwu/Clipsync/message"
	"github.com/google/uuid"
)

type Node struct {
	Id          uuid.UUID
	connections []string
	data        data.Data
	nodeType    int
}

func New() Node {
	id := uuid.New()

	return Node{
		Id:          id,
		connections: make([]string, 5),
	}
}

func (n *Node) HandleStream(conn net.Conn) {
	log.Printf("Node connected")

	defer conn.Close()
}

func (n *Node) handleMessage(message message.Message) {
	messageAsJson := message.ToJson()

	fmt.Println(messageAsJson)
}

func (n *Node) sendMessage(message_type int, receiver_id uuid.UUID) {
	message := message.New(n.Id, receiver_id, message_type)

	message.ToJson()
}

func (n *Node) AsServer() {
	n.nodeType = 0
	log.Printf("Server Node with id %s started , nodeType is %v", n.Id, n.nodeType)

	listener, err := net.Listen("tcp", "127.0.0.1:8080")

	n.handleError(err)

	for {
		if conn, err := listener.Accept(); err == nil {
			n.HandleStream(conn)
		}
	}
}

func (n *Node) AsClient() {
	n.nodeType = 1
	log.Printf("Client Node with id %s started , Node type is %v", n.Id, n.nodeType)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	n.handleError(err)

	defer conn.Close()

	n.HandleStream(conn)
}

func (n *Node) handleError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
