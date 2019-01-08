package brewery

import (
	"sync"

	rpio "github.com/stianeikeland/go-rpio"
)

type Controller interface {
	PowerPin(pin int, on bool) error
}

type GPIOController struct {
}

func PowerPin(pinNum int, on bool) error {
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
