package brewery

import (
	"context"
	"sync"

	model "github.com/golang001/brewery/model/gomodel"
)

type Brewery struct {
	scheme *model.ControlScheme
	mux    sync.RWMutex

	mashSensor  model.ThermometerClient
	boilSensor  model.ThermometerClient
	hermsSensor model.ThermometerClient

	element model.SwitchClient
}

func NewController() *Brewery {
	return &Brewery{}
}

func (c *Brewery) ReplaceConfig(scheme *model.ControlScheme) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.scheme = scheme
}

func (c *Brewery) Run() error {
	c.mux.RLock()
	defer c.mux.RUnlock()
	config := c.scheme
	switch sch := config.Scheme.(type) {
	case *model.ControlScheme_Boil_:
		_, err := c.element.On(context.Background(), &model.OnRequest{})
		return err
	case *model.ControlScheme_Mash_:
		res, err := c.boilSensor.Get(context.Background(), &model.GetRequest{})
		if err != nil {
			return err
		}
		if res.Temperature < sch.Mash.BoilMinTemp {
			_, err := c.element.On(context.Background(), &model.OnRequest{})
			if err != nil {
				return err
			}

		}
	case *model.ControlScheme_Power_:
	default:
		_, err := c.element.Off(context.Background(), &model.OffRequest{})
		return err
	}
	return nil
}
