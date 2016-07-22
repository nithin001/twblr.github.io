package control_structures

import "testing"

func TestCollatzChainLength(t *testing.T) {
	n := collatzChainLength(67)

	if n != 27 {
		t.Errorf("Expected %d, got %d", 27, n)
	}
}
