package detectable

import (
	"fmt"
	"math"
)

var InverseCSquared = 0.001

type Property struct {
	Intensity    float32
	PropertyType string
}

type PropertyType int

var (
	GammaRadiation PropertyType = 1
	EMRadiation    PropertyType = 2
	RFRadiation    PropertyType = 2
)

// TraceDetectable is a stuct point-in-time history of an object.
type TraceDetectable struct {
	Pos        Position
	Properties map[PropertyType]*Property
	ID         string
}

func NewTraceDetectable(x float64, y float64, z float64, t float64, id string) *TraceDetectable {
	return &TraceDetectable{Pos: Position{X: x, Y: y, Z: z, T: t}, ID: id,
		Properties: make(map[PropertyType]*Property)}
}

func (td *TraceDetectable) GetPosition() Position {
	return td.Pos
}

func (td *TraceDetectable) GetProperty(propType PropertyType) (*Property, error) {
	prop, ok := td.Properties[propType]
	if !ok {
		return nil, fmt.Errorf("property not found")
	}
	return prop, nil
}

func (td *TraceDetectable) GetID() string {
	return td.ID
}

func (td *TraceDetectable) Compare(req DetectRequest) bool {
	if req.Pos.T < td.Pos.T {
		return false
	}
	pos := td.GetPosition()
	deltaX := pos.X - req.Pos.X
	deltaY := pos.Y - req.Pos.Y
	deltaZ := pos.Z - req.Pos.Z
	distSquared := math.Pow(deltaX, 2) + math.Pow(deltaY, 2) + math.Pow(deltaZ, 2)
	lightTravelTimeSquared := distSquared * InverseCSquared
	deltaTimeSquared := math.Pow(req.Pos.T-pos.T, 2)

	return math.Abs(lightTravelTimeSquared-deltaTimeSquared) < 1
}

type DetectableHistory map[float64]Detectable

type DetectableDatabase struct {
	db map[string](*DetectableHistory)
}

func NewDetectableDatabase() *DetectableDatabase {
	return &DetectableDatabase{
		db: make(map[string]*DetectableHistory),
	}
}

func (db *DetectableDatabase) Register(obj Detectable) {
	hist, ok := db.db[obj.GetID()]
	if !ok {
		var tmpHist DetectableHistory = make(map[float64]Detectable)
		hist = &tmpHist
	}

	(*hist)[obj.GetPosition().T] = obj
	db.db[obj.GetID()] = hist
}

func (db *DetectableDatabase) Detect(req DetectRequest) DetectResponse {
	resp := DetectResponse{detected: []Detectable{}}
	for _, hist := range db.db {
		for i := req.Pos.T; i >= 0; i-- {
			d, ok := (*hist)[i]
			if !ok {
				continue
			}
			if d.Compare(req) {
				resp.detected = append(resp.detected, d)
				break
			}
		}
	}
	return resp
}
