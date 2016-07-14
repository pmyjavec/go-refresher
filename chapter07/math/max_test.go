package math

import "testing"

func TestMax(t *testing.T) {
	mx := max(1, 3, 10)
	if mx != 10 {
		t.Error("Expecting 10 got", mx)
	}
}
