package Studienarbeit_src

import "encoding/json"

type EventHandler func(*Event)

/*
Example Event:
{
	"event": "message",
	"data": "this is a data object"
}
*/

type Event struct {
	Name string      `json:"event"`
	Data interface{} `json:"data"`
}

func NewEventFromRaw(rawData []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(rawData, event)
	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
