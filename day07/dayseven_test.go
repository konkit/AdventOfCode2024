package dayseven

import (
	"testing"
)

func TestCheckEquations(t *testing.T) {
	lines := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	res := CheckEquations(lines)

	if res != 3749 {
		t.Errorf("Expected 3749, got %d", res)
	}
}

func TestCheckEquations2(t *testing.T) {
	lines := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	res := CheckEquations2(lines)

	if res != 11387 {
		t.Errorf("Expected 11387, got %d", res)
	}
}

func TestParseLine(t *testing.T) {
	line := "3267: 81 40 27"

	a, b := parseLine(line)

	if a != 3267 {
		t.Errorf("a: Expected 3267, got %d", a)
	}

	if b[0] != 81 {
		t.Errorf("b[0]: Expected 81, got %d", b[0])
	}

	if b[1] != 40 {
		t.Errorf("b[1]: Expected 40, got %d", b[1])
	}

	if b[2] != 27 {
		t.Errorf("b[2]: Expected 27, got %d", b[27])
	}
}
