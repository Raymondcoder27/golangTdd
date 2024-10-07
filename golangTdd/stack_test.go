package adt

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	s := NewStack() // No need for `adt.NewStack()` because we are in the same package.....

	if s.Empty() != true {
		t.Error("Stack was not empty")
	}
}
