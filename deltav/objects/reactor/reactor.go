package reactor

import (
	"github.com/golang001/deltav/objects/fuelTank"
)

type ReactorOutput struct {
	heat   int
	energy int
}

type Reactor interface {
	React(rate int) ReactorOutput
}

type FusionReactor struct {
	heliumBuffer    fuelTank.FuelTank
	deuteriumBuffer fuelTank.FuelTank
}
