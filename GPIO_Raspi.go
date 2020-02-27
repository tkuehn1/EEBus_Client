package Studienarbeit_src

import (
	"fmt"
	"github.com/stianeikeland/go-rpio" // you have to Import "go get github.com/stianeikeland/go-rpio for including the package for gpio on raspi "
	"time"
)

func Gpio() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(14)
	pin.Output()

	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
	pin.Write(rpio.Low)
}
