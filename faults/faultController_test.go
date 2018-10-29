package faults

import (
	"fmt"
	"testing"
)

func TesErrorFault(t *testing.T) {
	expectedErr := fmt.Errorf("error")
	errorFault := ErrorFault{Err: expectedErr, FailureRate: 4}

	err := errorFault.rateInvoke(int32(1))
	if err != nil {

	}
}
