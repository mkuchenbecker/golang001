package detectable

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TesDetectableLookback(t *testing.T) {
	InverseCSquared = 1 //Overwrite to make math easier.
	var pos PositionSystem = NewDetectableDatabase()
	pos.Register(NewTraceDetectable(0, 0, 0, 0, "A"))
	pos.Register(NewTraceDetectable(0, 0, 0, 1, "B"))
	pos.Register(NewTraceDetectable(0, 0, 0, 2, "C"))
	pos.Register(NewTraceDetectable(0, 0, 0, 3, "D"))
	pos.Register(NewTraceDetectable(0, 0, 0, 4, "E"))

	req := DetectRequest{Pos: Position{X: 1, Y: 0, Z: 0, T: 3},
		Range: 10, Filter: []PropertyType{GammaRadiation}}

	resp := pos.Detect(req)

	assert.Equal(t, 1, len(resp.detected))
	assert.Equal(t, "C", resp.detected[0].GetID())
}

func TestDetectableMultiItem(t *testing.T) {
	InverseCSquared = 1 //Overwrite to make math easier.
	var pos PositionSystem = NewDetectableDatabase()

	for i := float64(0); i < 100; i++ {
		r := rand.Float64() * 1000
		pos.Register(NewTraceDetectable(r, r, r, i, "A"))
		pos.Register(NewTraceDetectable(r, r, r, i, "B"))
		pos.Register(NewTraceDetectable(r, r, r, i, "C"))
		pos.Register(NewTraceDetectable(r, r, r, i, "D"))
		pos.Register(NewTraceDetectable(r, r, r, i, "E"))
		pos.Register(NewTraceDetectable(r, r, r, i, "F"))
	}

	pos.Register(NewTraceDetectable(2, 0, 0, 10, "A"))
	pos.Register(NewTraceDetectable(0, 2, 0, 10, "B"))
	pos.Register(NewTraceDetectable(0, 0, 2, 10, "C"))
	pos.Register(NewTraceDetectable(-2, 0, 0, 10, "D"))
	pos.Register(NewTraceDetectable(0, -2, 0, 10, "E"))
	pos.Register(NewTraceDetectable(0, 0, -2, 10, "F"))

	assert.Equal(t, 6, len(pos.(*DetectableDatabase).db))
	assert.Equal(t, 100, len(*(pos.(*DetectableDatabase).db["A"])))
	assert.Equal(t, 100, len(*(pos.(*DetectableDatabase).db["F"])))

	req := DetectRequest{Pos: Position{X: 0, Y: 0, Z: 0, T: 12},
		Range: 10, Filter: []PropertyType{GammaRadiation}}

	resp := pos.Detect(req)

	assert.Equal(t, 6, len(resp.detected))
	countMap := make(map[string]int)
	for _, val := range resp.detected {
		// Each key should be unique A-F
		_, ok := countMap[val.GetID()]
		assert.False(t, ok)
		countMap[val.GetID()]++
	}
}

func TestPerformance(t *testing.T) {
	InverseCSquared = 0.001 //Overwrite to make math easier.
	total := 1000.0         //* 1000.0
	objs := []*TraceDetectable{}
	for i := float64(0); i < total; i++ {
		objs = append(objs,
			NewTraceDetectable(0, 0, 0, 0, fmt.Sprintf("%f", i)))
	}

	var pos PositionSystem = NewDetectableDatabase()

	gameDurationMinutes := float64(60)

	for i := float64(0); i <= gameDurationMinutes*60; i++ {
		for _, obj := range objs {
			obj.Pos.X += (rand.Float64())
			obj.Pos.Y += (rand.Float64())
			obj.Pos.Z += (rand.Float64())
			obj.Pos.T = i
			pos.Register(obj)
		}
	}

	beforeTime := time.Now()
	req := DetectRequest{Pos: Position{X: 0, Y: 0, Z: 0, T: 50},
		Range: 10, Filter: []PropertyType{GammaRadiation}}

	pos.Detect(req)
	afterTime := time.Now()

	assert.WithinDuration(t, beforeTime, afterTime, 1000*time.Millisecond)

}
