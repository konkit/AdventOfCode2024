package dayeight

import "testing"

func TestCalculateAntinodes(t *testing.T) {
	lines := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	res := CalculateAntinodes(lines)

	if res != 14 {
		t.Errorf("Expected 14, got %d", res)
	}
}

func TestCalculateAntinodes2(t *testing.T) {
	lines := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	res := CalculateAntinodes2(lines)

	if res != 34 {
		t.Errorf("Expected 34, got %d", res)
	}
}
