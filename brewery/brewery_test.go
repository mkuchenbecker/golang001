package brewery

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	model "github.com/golang001/brewery/model/gomodel"

	"github.com/golang/mock/gomock"
	mocks "github.com/golang001/brewery/model/gomock"
)

func TestElementPowerLevelToggle(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)

	// Base test.
	mockSwitch.EXPECT().On(context.Background(), gomock.Any()).Return(&model.OnResponse{}, nil).Times(1)
	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	// Interval is 1 second, time on is 0.75s, quit after 1 second.
	// Result is that the switch should turn on, .75s later turn off, and on the next iteration quit.

	brewery := Brewery{element: mockSwitch}
	err := brewery.elementPowerLevelToggle(
		time.Duration(750)*time.Millisecond,
		time.Duration(1)*time.Second,
		time.Duration(1)*time.Second)
	assert.NoError(t, err)
}
func TestElementPowerLevelToggleMultipleLoops(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)
	// Tests we always turn off after every
	mockSwitch.EXPECT().On(context.Background(), gomock.Any()).Return(&model.OnResponse{}, nil).Times(10)
	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(10)
	// Interval is 1 second, time on is 0.75s, quit after 1.25 second.
	// Two iterations should occur.
	brewery := Brewery{element: mockSwitch}
	err := brewery.elementPowerLevelToggle(
		time.Duration(50)*time.Millisecond,
		time.Duration(1000)*time.Millisecond,
		time.Duration(100)*time.Millisecond)
	assert.NoError(t, err)
}

func TestElementOnError(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)
	mockSwitch.EXPECT().On(context.Background(),
		gomock.Any()).Return(&model.OnResponse{}, fmt.Errorf("unable to turn coil on")).Times(1)
	mockSwitch.EXPECT().Off(context.Background(),
		gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	brewery := Brewery{element: mockSwitch}
	err := brewery.ElementOn(1 * time.Millisecond)
	assert.Error(t, err)

	mockSwitch.EXPECT().On(context.Background(),
		gomock.Any()).Return(&model.OnResponse{}, nil).Times(1)
	// We always attempt to turn the coiloff 3 times.
	mockSwitch.EXPECT().Off(context.Background(),
		gomock.Any()).Return(&model.OffResponse{}, fmt.Errorf("unable to turn coil off")).Times(3)
	err = brewery.ElementOn(1 * time.Millisecond)
	assert.Error(t, err)
}

func TestElementPowerLevelToggleError(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)
	mockSwitch.EXPECT().On(context.Background(),
		gomock.Any()).Return(&model.OnResponse{}, fmt.Errorf("unable to turn coil on")).Times(1)
	mockSwitch.EXPECT().Off(context.Background(),
		gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	brewery := Brewery{element: mockSwitch}
	err := brewery.elementPowerLevelToggle(9*time.Millisecond, 100*time.Millisecond, 10*time.Millisecond)
	assert.Error(t, err)

	mockSwitch.EXPECT().On(context.Background(),
		gomock.Any()).Return(&model.OnResponse{}, nil).Times(1)
	mockSwitch.EXPECT().Off(context.Background(),
		gomock.Any()).Return(&model.OffResponse{}, fmt.Errorf("unable to turn coil off")).Times(3)

	err = brewery.elementPowerLevelToggle(9*time.Millisecond, 100*time.Millisecond, 10*time.Millisecond)
	assert.Error(t, err)

}

func TestElementPowerLevel0(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)

	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	brewery := Brewery{element: mockSwitch}
	err := brewery.ElementPowerLevel(0, 1)
	assert.NoError(t, err)
}

func TestElementPowerLevel101(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)

	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	brewery := Brewery{element: mockSwitch}
	err := brewery.ElementPowerLevel(101, 1)
	assert.NoError(t, err)
}

func TestElementPowerLevel100(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSwitch := mocks.NewMockSwitchClient(mockCtrl)

	mockSwitch.EXPECT().Off(context.Background(), gomock.Any()).Return(&model.OffResponse{}, nil).Times(1)
	mockSwitch.EXPECT().On(context.Background(), gomock.Any()).Return(&model.OnResponse{}, nil).Times(1)
	brewery := Brewery{element: mockSwitch}
	err := brewery.ElementPowerLevel(100, 1)
	assert.NoError(t, err)
}
