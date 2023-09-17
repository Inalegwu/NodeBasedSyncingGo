package message

import (
	"encoding/json"
	"fmt"
)

const (
	Connect           = 0
	AcceptConnection  = 1
	DeclineConnection = 2
	Sync              = 3
)

type Message struct {
	id          string
	sender_id   string
	receiver_id string
	messageType int
}

func New(sender_id string, receiver_id string, message_type int) Message {
	return Message{
		id:          "1",
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
