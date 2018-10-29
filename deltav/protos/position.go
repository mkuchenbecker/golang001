package deltav_protos

import "math"

func (p *Vector3) MagnitudeSquared() float64 {
	return math.Pow(p.X, 2) + math.Pow(p.Y, 2) + math.Pow(p.Z, 2)
}

func (p *Vector3) Subtract(n *Vector3) *Vector3 {
	return &Position{X: p.X - n.X, Y: p.Y - n.Y, Z: p.Z - n.ZT}
}

func (p *Position) Subtract(n *Position) *Position {
	return &Position{Position: p.Position.Subtract(n.position), T: p.T - n.T}
}
