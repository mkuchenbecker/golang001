package brewery

import (
	"context"
	"fmt"
	"sync"
	"time"

	model "github.com/golang001/brewery/model/gomodel"
)

func print(s string) {
	fmt.Printf("%s %s\n", time.Now().Format(time.StampMilli), s)
}

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
	if resHerms.Temperature > c.scheme.GetMash().HermsMaxTemp { //Don't want to overshoot.
		return false, nil
	}

	resMash, err := c.mashSensor.Get(context.Background(), &model.GetRequest{})
	if err != nil {
		return false, err
	}
	if resMash.Temperature < c.scheme.GetMash().MashMinTemp {
		on = true
		return
	}
	return false, nil
}

func (c *Brewery) ElementOff() error {
	var err error
	for i := 0; i < 3; i++ {
		_, err = c.element.Off(context.Background(), &model.OffRequest{})
		if err == nil {
			return err
		}
	}
	return err
}

func (b *Brewery) Run(ttlSec int) error {
	b.mux.RLock()
	defer b.mux.RUnlock()
	ttl := time.Duration(ttlSec) * time.Second
	config := b.scheme
	switch sch := config.Scheme.(type) {
	case *model.ControlScheme_Boil_:
		err := b.ElementOn(ttl)
		return err
	case *model.ControlScheme_Mash_:
		on, err := b.mashThermOn()
		if err != nil {
			return err
		}
		if !on {
			err := b.ElementOff()
			if err != nil {
				return err
			}
		}
		b.ElementOn(ttl)
		return err
	case *model.ControlScheme_Power_:
		b.ElementPowerLevel(int(sch.Power.PowerLevel), ttlSec) // Toggle for one hour.
	default:
	}
	return b.ElementOff()

}

func (b *Brewery) ElementOn(ttl time.Duration) (err error) {
	defer func() {
		offErr := b.ElementOff()
		if offErr != nil || err != nil {
			err = fmt.Errorf("errors occured: '%s', '%s", offErr, err)
		}
	}()

	print("Turning coil on.")
	_, err = b.element.On(context.Background(), &model.OnRequest{})
	if err != nil {
		print(fmt.Sprintf("encountered error turning coil on: %+v", err))
		return err
	}
	print("Sleep.")
	timer := time.NewTimer(ttl)
	<-timer.C
	return err
}

func (b *Brewery) ElementPowerLevel(powerLevel int, ttlSeconds int) error {
	ttl := time.Duration(ttlSeconds) * time.Second
	if powerLevel < 1 {
		err := b.ElementOff()
		if err != nil {
			return err
		}
	}
	if powerLevel > 100 {
		err := b.ElementOff()
		if err != nil {
			return err
		}
	}
	if powerLevel == 100 {
		err := b.ElementOn(ttl)
		if err != nil {
			return err
		}
	}
	interval := 2
	delay := time.Duration(powerLevel / 100 * interval)
	return b.elementPowerLevelToggle(delay, ttl, time.Duration(interval)*time.Second)
}

func (b *Brewery) elementPowerLevelToggle(delay time.Duration, ttl time.Duration, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	quit := make(chan bool)
	resErr := make(chan error)

	go func() {
		for {
			print("loop")
			select {
			case <-ticker.C:
				print("tick")
				err := b.ElementOn(delay)
				if err != nil {
					resErr <- err
					return
				}
				print("tock")
			case <-quit:
				ticker.Stop()
				resErr <- nil
				return
			}
		}
	}()

	go func() { // Make sure the process always exits.
		print("waiting")
		timer := time.NewTimer(ttl)
		<-timer.C
		print("quit")
		quit <- true
	}()

	return <-resErr
}
