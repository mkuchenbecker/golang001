package mastercontrol

import "github.com/golang001/deltav/detectable"

type MasterControl struct {
	database detectable.DetectableDatabase

	//Map of ids to their objects.'s.
	objects map[string]interface{}
}
