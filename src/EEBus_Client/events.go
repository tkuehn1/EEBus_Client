package EEBus_Client

import (
	"encoding/json"
	"log"
)

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
	Pinnumber int  `json:"pinnumber"`
	Active    bool `json:"active"`
	Status    bool `json:"status"` //work in progress
}

func NewEventFromRaw(rawData []byte) (*Event, error) {
	log.Printf("%s\n", rawData)
	event := new(Event)
	err := json.Unmarshal(rawData, event)
	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}

func handleEventMessage(d Data) error {
	err := Gpio(d.Pinnumber, d.Active)
	if err != nil {
		return err
	}
	return nil
}
