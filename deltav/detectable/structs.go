package detectable

import (
	"math"

	"github.com/golang001/deltav/common"
)

const C = float64(1000) // Units / second
var InverseCSquared = math.Pow(1/C, 2)

func DefaultInverseC() float64 {
	return math.Pow(float64(1)/C, 2)
}

type Property struct {
	Intensity    float32
	PropertyType PropertyType
}

type PropertyType int

var (
	GammaRadiation PropertyType = 1
	EMRadiation    PropertyType = 2
	RFRadiation    PropertyType = 2
)

type DetectRequest struct {
	Pos    common.Position
	Range  float64
	Filter []PropertyType
}

type DetectResponse struct {
	detected []Detectable
}

type PositionSystem interface {
	Detect(DetectRequest) DetectResponse
	Register(Detectable)
}

type Detectable interface {
	GetPosition() common.Position
	GetProperty(PropertyType) (*Property, error)
	GetID() string
	Compare(DetectRequest) bool
}
