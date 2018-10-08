package detectable

import "math"

type Position struct {
	X float64
	Y float64
	Z float64
	T float64
}

func (p *Position) MagnitudeSquared() float64 {
	return math.Pow(p.X, 2) + math.Pow(p.Y, 2) + math.Pow(p.Z, 2)
}

// type NamedObject struct {
// 	GetName() string
// }

type DetectRequest struct {
	Pos    Position
	Range  int
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
