package fuelTank

type FuelTank interface {
	// AddFuel adds fuel and returns the volume of unused fuel.
	AddFuel(volume int) int
	WithdrawFuel(volume int) int
	GetVolume() int
	GetCapacity() int
}

type GasFuelTank struct {
	capacity int
	current  int
}

func NewGasTank(capacity int, initialVolume int) *GasFuelTank {
	return &GasFuelTank{capacity: capacity, current: initialVolume}
}

func (gt *GasFuelTank) AddFuel(volume int) int {
	spareCap := gt.capacity - volume + gt.current
	if spareCap < 0 {
		gt.current = gt.capacity
		return volume + spareCap
	}
	return 0
}

func (gt *GasFuelTank) WithdrawFuel(volume int) int {
	if gt.current < volume {
		gt.current = 0
		return gt.current
	}
	gt.current -= volume
	return volume
}

func (gt *GasFuelTank) GetVolume() int {
	return gt.current
}

func (gt *GasFuelTank) GetCapacity() int {
	return gt.capacity
}
