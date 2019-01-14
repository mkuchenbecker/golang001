package brewery

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/golang001/brewery/mocks"
)

func TestReadTemperature(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockController := mocks.NewController(mockCtrl)
	mockController.EXPECT().ReadTemperature(TemperatureAddress("address123")).Return(float64(50), nil).Times(1)
}
