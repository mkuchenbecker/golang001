package deltav_model

/*
This file is a value-added addition helper functions on the position/vector struct.
*/
import "math"

func (p *Vector3) MagnitudeSquared() float64 {
	return math.Pow(p.X, 2) + math.Pow(p.Y, 2) + math.Pow(p.Z, 2)
}

func (p *Vector3) Subtract(n *Vector3) *Vector3 {
	return &Vector3{X: p.X - n.X, Y: p.Y - n.Y, Z: p.Z - n.Z}
}

func (p *Position) Subtract(n *Position) *Position {
	return &Position{Position: p.Position.Subtract(n.Position), T: p.T - n.T}
}
