package brewery

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	model "github.com/golang001/brewery/model/gomodel"

	"github.com/golang/mock/gomock"
	mocks "github.com/golang001/brewery/model/gomock"
)

func TestTimerToggle(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)

	// Base test.
	mockSwitch.EXPECT().On(context.Background(), gomock.Any()).Return(&model.OnResponse{}, nil).Times(1)
	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	// Interval is 1 second, time on is 0.75s, quit after 1 second.
	// Result is that the switch should turn on, .75s later turn off, and on the next iteration quit.
	err := toggle(mockSwitch,
		time.Duration(750)*time.Millisecond,
		time.Duration(1)*time.Second,
		time.Duration(1)*time.Second)
	assert.NoError(t, err)
}
func TestTimerToggleMultipleCalls(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)
	// Tests we always turn off after every
	mockSwitch.EXPECT().On(context.Background(), gomock.Any()).Return(&model.OnResponse{}, nil).Times(10)
	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(10)
	// Interval is 1 second, time on is 0.75s, quit after 1.25 second.
	// Two iterations should occur.
	err := toggle(mockSwitch,
		time.Duration(50)*time.Millisecond,
		time.Duration(1000)*time.Millisecond,
		time.Duration(100)*time.Millisecond)
	assert.NoError(t, err)
}

func TestTimerTogglePowerEdges(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)

	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	err := ToggleSwitch(mockSwitch, 0, 1)
	assert.NoError(t, err)

	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	err = ToggleSwitch(mockSwitch, 101, 1)
	assert.NoError(t, err)

	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	mockSwitch.EXPECT().On(context.Background(), gomock.Any()).Return(&model.OnResponse{}, nil).Times(1)
	err = ToggleSwitch(mockSwitch, 100, 1)
	assert.NoError(t, err)
}
