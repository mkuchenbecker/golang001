package detectable

import "math"

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

type Position struct {
	X float64
	Y float64
	Z float64
	T float64
}

func (p *Position) MagnitudeSquared() float64 {
	return math.Pow(p.X, 2) + math.Pow(p.Y, 2) + math.Pow(p.Z, 2)
}

func (p *Position) Subtract(n *Position) *Position {
	return &Position{X: p.X - n.X, Y: p.Y - n.Y, Z: p.Z - n.Z, T: p.T - n.T}
}

type DetectRequest struct {
	Pos    Position
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
	GetPosition() Position
	GetProperty(PropertyType) (*Property, error)
	GetID() string
	Compare(DetectRequest) bool
}
