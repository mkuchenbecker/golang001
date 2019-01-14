package sensors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/golang001/brewery/model/gomodel"
	"github.com/golang001/brewery/rpi/gpio"

	"github.com/golang/mock/gomock"
	mocks "github.com/golang001/brewery/rpi/gpio/mocks"
)

func TestReadTemperature(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := mocks.NewMockController(mockCtrl)
	mockController.EXPECT().ReadTemperature(gpio.TemperatureAddress("address123")).Return(float64(50), nil).Times(1)

	server := ThermometerServer{ctrl: mockController, address: gpio.TemperatureAddress("address123")}

	res, err := server.Get(context.Background(), &model.GetRequest{})
	assert.NoError(t, err)
	assert.Equal(t, &model.GetResponse{Temperature: 50}, res)
}
