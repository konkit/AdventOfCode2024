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

func TestCalculatePrice2(t *testing.T) {
	lines := []string{
		"AAAAAA",
		"AAABBA",
		"AAABBA",
		"ABBAAA",
		"ABBAAA",
		"AAAAAA",
	}

	res := CalculatePrice2(lines)

	if res != 368 {
		t.Errorf("got %d, want 368", res)
	}
}

func TestCalculatePrice22(t *testing.T) {
	lines := []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}

	res := CalculatePrice2(lines)

	if res != 236 {
		t.Errorf("got %d, want 236", res)
	}
}
