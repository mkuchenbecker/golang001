package brewery

import (
	"fmt"
	"sync"

	temperature "github.com/yryz/ds18b20"

	rpio "github.com/stianeikeland/go-rpio"
)

type TemperatureAddress string

func NewTemperatureAddress(address string) (TemperatureAddress, error) {
	sensors, err := temperature.Sensors()
	if err != nil {
		return "", err
	}

	for _, sensor := range sensors {
		if sensor == address {
			return TemperatureAddress(address), nil
		}
	}

	return "", fmt.Errorf("sensor not found %s", address)
}

type Controller interface {
	PowerPin(pin int, on bool) error
	ReadTemperature(address TemperatureAddress) (float64, error)
}

type GPIOController struct {
}

func (gpio *GPIOController) PowerPin(pinNum int, on bool) error {
	err := rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()
	pin := rpio.Pin(pinNum)
	if on {
		pin.High()
	} else {
		pin.Low()
	}
	return nil
}

func (gp *GPIOController) ReadTemperature(address TemperatureAddress) (float64, error) {
	return temperature.Temperature(string(address))
}

type FakeController struct {
	mux  sync.Mutex
	pins map[int]bool
	err  error
}

func (fc *FakeController) PowerPin(pinNum int, on bool) error {
	fc.mux.Lock()
	defer fc.mux.Unlock()
	fc.pins[pinNum] = on
	return fc.err
}
