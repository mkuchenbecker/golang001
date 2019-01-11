package brewery

import (
	"fmt"
	"sync"

	temperature "github.com/yryz/ds18b20"

	rpio "github.com/stianeikeland/go-rpio"
)

type TemperatureAddress string

type SensorArray interface {
	Sensors() ([]Sensor, error)
	Temperature(Sensor) (Celsius, error)
}

type DefaultSensorArray struct {
}

type Sensor string
type Celsius float64

func (d *DefaultSensorArray) Sensors() ([]Sensor, error) {
	strSli, err := temperature.Sensors()
	retSli := make([]Sensor, 0)
	if err != nil {
		return retSli, err
	}
	for _, s := range strSli {
		retSli = append(retSli, Sensor(s))
	}
	return retSli, nil
}

func (d *DefaultSensorArray) Temperature(sensor Sensor) (Celsius, error) {
	temp, err := temperature.Temperature(string(sensor))
	return Celsius(temp), err
}

func NewTemperatureAddress(address string, sensorArray SensorArray) (TemperatureAddress, error) {
	sensors, err := sensorArray.Sensors()
	if err != nil {
		return "", err
	}

	for _, sensor := range sensors {
		if sensor == Sensor(address) {
			return TemperatureAddress(address), nil
		}
	}

	return "", fmt.Errorf("sensor not found %s", address)
}

type Controller interface {
	PowerPin(pin int, on bool) error
	ReadTemperature(address TemperatureAddress) (float64, error)
}

type GPIOPins {
	Open() error
	Close() error //more needed, check retuen types
}

type GPIOController struct {
	sensorArray SensorArray
	gpioPins	GPIOPins
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

func (gp *GPIOController) ReadTemperature(sensor Sensor) (Celsius, error) {
	return gp.sensorArray.Temperature(sensor)
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
