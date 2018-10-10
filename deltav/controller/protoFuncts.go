package controller

import "math"

func (p *Position) MagnitudeSquared() float64 {
	return math.Pow(p.X, 2) + math.Pow(p.Y, 2) + math.Pow(p.Z, 2)
}

func (p *Position) Subtract(n *Position) *Position {
	return &Position{X: p.X - n.X, Y: p.Y - n.Y, Z: p.Z - n.Z, T: p.T - n.T}
}
