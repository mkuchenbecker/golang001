package gpio

type FakeGPIOPin struct {
	pin   uint8
	calls []string
}

func NewFakeGPIOPin(pin uint8) *FakeGPIOPin {
	return &FakeGPIOPin{pin: pin, calls: make([]string, 0)}
}

func (p *FakeGPIOPin) High() {
	p.calls = append(p.calls, "high")
}

func (p *FakeGPIOPin) Low() {
	p.calls = append(p.calls, "low")
}
