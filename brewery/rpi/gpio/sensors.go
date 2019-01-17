//+build !test

package gpio

import (
	"fmt"

	"github.com/golang001/brewery/rpi/gpio/igpio"
	temperature "github.com/yryz/ds18b20"
)

type DefaultSensorArray struct {
}

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

func NewTemperatureAddress(address string, sensorArray SensorArray) (igpio.TemperatureAddress, error) {
	sensors, err := sensorArray.Sensors()
	if err != nil {
		return "", err
	}

	for _, sensor := range sensors {
		if sensor == Sensor(address) {
			return igpio.TemperatureAddress(address), nil
		}
	}

	return "", fmt.Errorf("sensor not found %s", address)
}
