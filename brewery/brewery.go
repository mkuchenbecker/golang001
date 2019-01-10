package brewery

import (
	"context"
	"fmt"
	"sync"
	"time"

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

func NewBrewery() *Brewery {
	return &Brewery{}
}

func (c *Brewery) ReplaceConfig(scheme *model.ControlScheme) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.scheme = scheme
}

func (c *Brewery) mashThermOn() (on bool, err error) {
	resBoil, err := c.boilSensor.Get(context.Background(), &model.GetRequest{})
	if err != nil {
		return false, err
	}
	if resBoil.Temperature < c.scheme.GetMash().BoilMinTemp {
		on = true
		return
	}

	resHerms, err := c.hermsSensor.Get(context.Background(), &model.GetRequest{})
	if err != nil {
		return false, err
	}
	if resHerms.Temperature < c.scheme.GetMash().HermsMinTemp {
		on = true
		return
	}

	resMash, err := c.mashSensor.Get(context.Background(), &model.GetRequest{})
	if err != nil {
		return false, err
	}
	if resMash.Temperature < c.scheme.GetMash().MashMinTemp && resHerms.Temperature-resMash.Temperature < 5.0 {
		on = true
		return
	}
	return false, nil
}

func (c *Brewery) CoilOff() {
	var err error
	for i := 0; i < 3; i++ {
		_, err = c.element.Off(context.Background(), &model.OffRequest{})
		if err == nil {
			return
		}
	}
	if err != nil {
		panic("unable to turn off coil")
	}
}

func (c *Brewery) Run(ttlSec int) error {
	c.mux.RLock()
	defer c.mux.RUnlock()
	config := c.scheme
	switch sch := config.Scheme.(type) {
	case *model.ControlScheme_Boil_:
		_, err := c.element.On(context.Background(), &model.OnRequest{})
		return err
	case *model.ControlScheme_Mash_:
		on, err := c.mashThermOn()
		if err != nil {
			return err
		}
		if !on {
			c.CoilOff()
		}
		TurnOnCoil(c.element, time.Duration(ttlSec)*time.Second)
		return err
	case *model.ControlScheme_Power_:
		ToggleSwitch(c.element, int(sch.Power.PowerLevel), ttlSec) // Toggle for one hour.
	default:
	}
	_, err := c.element.Off(context.Background(), &model.OffRequest{})
	return err

}

func TurnOnCoil(s model.SwitchClient, ttl time.Duration) (err error) {
	defer func() { //Turning off is deferred so it always runs.
		fmt.Println("Turning coil off.")
		_, offErr := s.Off(context.Background(), &model.OffRequest{})
		if offErr != nil || err != nil {
			err = fmt.Errorf("encountered heating element errors: '%+v', '%+v'", err, offErr)
		}
	}()

	fmt.Println("Turning coil on.")
	_, err = s.On(context.Background(), &model.OnRequest{})
	if err != nil {
		fmt.Printf("encountered error turning coil on: %+v", err)
		return err
	}
	fmt.Println("Sleep.")
	timer := time.NewTimer(ttl)
	<-timer.C
	return err
}

func ToggleSwitch(s model.SwitchClient, powerLevel int, ttlSeconds int) error {
	ttl := time.Duration(ttlSeconds) * time.Second
	if powerLevel < 1 {
		_, err := s.Off(context.Background(), &model.OffRequest{})
		return err
	}
	if powerLevel > 100 {
		_, err := s.Off(context.Background(), &model.OffRequest{})
		return err
	}
	if powerLevel == 100 {
		TurnOnCoil(s, ttl)
	}
	interval := 2
	delay := time.Duration(powerLevel / 100 * interval)
	return toggle(s, delay, ttl, time.Duration(interval)*time.Second)
}

func toggle(s model.SwitchClient, delay time.Duration, ttl time.Duration, interval time.Duration) (err error) {
	ticker := time.NewTicker(interval)
	quit := make(chan bool)
	defer func() { // Make sure the process always exits.
		quit <- true
	}()
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("tick")
				err = TurnOnCoil(s, delay)
				if err != nil {
					ticker.Stop()
					return
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	if err != nil {
		return err
	}

	timer := time.NewTimer(ttl)
	<-timer.C

	return
}
