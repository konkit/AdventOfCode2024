package day12

import "testing"

func TestCalculatePrice(t *testing.T) {
	lines := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}

	res := CalculatePrice(lines)

	if res != 140 {
		t.Errorf("got %d, want 140", res)
	}
}
