package gpio

type TemperatureAddress string

type SensorArray interface {
	Sensors() ([]Sensor, error)
	Temperature(Sensor) (Celsius, error)
}

type Sensor string
type Celsius float64

type Controller interface {
	PowerPin(pin int, on bool) error
	ReadTemperature(address TemperatureAddress) (float64, error)
}

type IGpio interface {
	Open() error
	Close() error //more needed, check retuen types
	Pin(uint8) PinGpio
}

type PinGpio interface {
	High()
	Low()
}
