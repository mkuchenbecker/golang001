package faults

import (
	"math/rand"
	"time"
)

// A fault is a function that can inject random failures into.
type Fault interface {
	Invoke() error
}

// ErrorFault is a function which returns an error with a given probability.
type ErrorFault struct {
	Err         error
	FailureRate int32
	counter     int64
}

func (ef *ErrorFault) rateInvoke(rate int32) error {
	if rand.Int31()%rate == 0 {
		ef.counter++
		return ef.Err
	}
	return nil
}

// Invoke may return an error.
func (ef *ErrorFault) Invoke() error {
	return ef.rateInvoke(ef.FailureRate)
}

// LatencyFault is a function that randomly injects latency into a function.
type LatencyFault struct {
	Latency     time.Duration
	FailureRate int32
	counter     int64
}

func (lf *LatencyFault) rateInvoke(rate int32) {

	if rand.Int31()%lf.FailureRate == 0 {
		time.Sleep(lf.Latency)
		lf.counter++
	}
}

// Invoke may inject latency.
func (lf *LatencyFault) Invoke() error {
	lf.rateInvoke(lf.FailureRate)
	return nil
}

// FaultInjector is a collection of faults that can be applied to a function.
type FaultInjector struct { // 1/X chance of failure.
	Faults []Fault
}

type FallibleFn func(...interface{}) (interface{}, error)

func (fi *FaultInjector) Inject(fn FallibleFn, args ...interface{}) (interface{}, error) {
	for _, fault := range fi.Faults {
		err := fault.Invoke()
		if err != nil {
			return nil, err
		}
	}
	return fn(args...)
}
