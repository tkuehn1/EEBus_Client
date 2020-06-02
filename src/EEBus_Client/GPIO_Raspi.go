package EEBus_Client

import (
	"fmt"
	"github.com/stianeikeland/go-rpio" // you have to Import "go get github.com/stianeikeland/go-rpio for including the package for gpio on raspi "
)

func Gpio(pinnumber int, active bool) error {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		return err
	}

	defer rpio.Close()

	pin := rpio.Pin(pinnumber)
	pin.Output()

	if active {
		pin.High()
	} else {
		pin.Low()
	}
	return nil
}

func Taster() string {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()
	pin := rpio.Pin(14)
	pin.Input()
	for {
		res := pin.Read()

		if res == rpio.High {
			return "high"
		} else if res == rpio.Low {
			return "low"
		}
	}
}

func Led(status string) error {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		return err
	}

	defer rpio.Close()

	pin := rpio.Pin(14)
	pin.Output()

	if status == "high" {
		pin.High()
	} else if status == "low" {
		pin.Low()
	} else {
		return err
	}
	return nil
}
