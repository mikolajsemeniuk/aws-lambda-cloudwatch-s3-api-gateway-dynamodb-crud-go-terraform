package api

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Key       string
	Name      string
	Amount    int
	Price     float32
	Completed bool
	Created   string
	Updated   string
}

type OrderInput struct {
	Name      string
	Amount    int
	Price     float32
	Completed bool
}

type Message struct {
	Message any
}

func (m Message) MarshalJSON() ([]byte, error) {
	switch v := m.Message.(type) {
	case error:
		if v != nil {
			return json.Marshal(v.Error())
		}
		return json.Marshal("")
	case string:
		return json.Marshal(v)
	default:
		return nil, fmt.Errorf("unsupported type: %T", m.Message)
	}
}
