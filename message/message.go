package message

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

const (
	Connect           = 0
	AcceptConnection  = 1
	DeclineConnection = 2
	Sync              = 3
)

type Message struct {
	id          uuid.UUID
	sender_id   uuid.UUID
	receiver_id uuid.UUID
	messageType int
}

func New(sender_id uuid.UUID, receiver_id uuid.UUID, message_type int) Message {
	return Message{
		id:          uuid.New(),
		sender_id:   sender_id,
		receiver_id: receiver_id,
		messageType: message_type,
	}
}

func (m *Message) ToJson() []byte {
	asJson, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Unable to marshal message to json")
	}

	return asJson
}
