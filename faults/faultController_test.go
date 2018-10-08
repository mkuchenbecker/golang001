package faults

import (
	"fmt"
	"testing"
)

func TestErrorFault(t *testing.T) {
	expectedErr := fmt.Errorf("error")
	errorFault := ErrorFault{Err: expectedErr, FailureRate: 4}

	err := errorFault.rateInvoke()
}
