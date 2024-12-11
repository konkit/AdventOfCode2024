package dayten

import "testing"

func TestCalculate(t *testing.T) {
	input := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	res := Calculate(input)

	if res != 36 {
		t.Error("Expected 36, got ", res)
	}
}

func TestCalculateSmaller(t *testing.T) {
	input := []string{
		"...0...",
		"...1...",
		"...2...",
		"6543456",
		"7.....7",
		"8.....8",
		"9.....9",
	}

	res := Calculate(input)

	if res != 2 {
		t.Error("Expected 2, got ", res)
	}
}

func TestCalculateSmaller2(t *testing.T) {
	input := []string{
		"..90..9",
		"...1.98",
		"...2..7",
		"6543456",
		"765.987",
		"876....",
		"987....",
	}

	res := Calculate(input)

	if res != 4 {
		t.Error("Expected 4, got ", res)
	}
}
