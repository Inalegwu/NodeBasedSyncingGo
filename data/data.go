package data

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	id      string
	content string
}

func New(content string) Data {
	return Data{
		id:      "1",
		content: content,
	}
}

func (d *Data) ToJson() []byte {
	dataAsJson, err := json.Marshal(d)
	if err != nil {
		fmt.Println("Error occurred while marshalling json ")
	}
	return dataAsJson
}
