package reactor

import (
	"testing"

	model "github.com/golang001/deltav/model/gomodel"
	"github.com/stretchr/testify/assert"
)

func TestReactor(t *testing.T) {
	reactor := FusionReactor{}
	storageRequest := model.StorageRequest{}
	assert.NotEqual(t, reactor, storageRequest)

	// Figure out grpc mocks
}
