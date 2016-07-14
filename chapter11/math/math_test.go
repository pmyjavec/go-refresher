package math

import "testing"

func TestAvage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})

	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}

	v = Average([]float64{})

	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}
