package gpio

import (
	"sync"

	rpio "github.com/stianeikeland/go-rpio"
)

type GPIOController struct {
	sensorArray SensorArray
	gpioPins    IGpio
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
