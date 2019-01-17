package gpio

import (
	"github.com/golang001/brewery/rpi/gpio/igpio"
	"github.com/golang001/brewery/utils"
)

type GPIOController struct {
	sensors  igpio.SensorArray
	gpioPins igpio.IGpio
}

func (gpio *GPIOController) PowerPin(pinNum uint8, on bool) (err error) {
	err = gpio.gpioPins.Open()
	if err != nil {
		return err
	}
	defer utils.DeferErrReturn(gpio.gpioPins.Close, &err)
	pin := gpio.gpioPins.Pin(pinNum)
	if on {
		pin.High()
	} else {
		pin.Low()
	}
	return nil
}

func (gp *GPIOController) ReadTemperature(sensor igpio.Sensor) (igpio.Celsius, error) {
	return gp.sensors.Temperature(sensor)
}
