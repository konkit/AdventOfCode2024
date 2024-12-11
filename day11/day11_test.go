package day11

import "testing"

func TestCalculate(t *testing.T) {
	line := "125 17"

	res := Calculate(line)

	if res != 55312 {
		t.Errorf("Expected 55312, got %d", res)
	}
}
