package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	assert.Equal(t, "1", "1")
}
