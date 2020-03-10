package Studienarbeit_src

import "encoding/json"

type EventHandler func(*Event)

/*
Example Event:
{
	"event": "message",
	"data": {
		"pinnumber": 16,
		"active": true,
		"status": false,
	}
}
*/

type Event struct {
	Name string `json:"event"`
	Data Data   `json:"data"`
}

type Data struct {
	Pinnumber float64 `json:"pinnumber"`
	Active    bool    `json:"active"`
	Status    bool    `json:"status"`
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
