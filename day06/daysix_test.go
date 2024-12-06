package day06

import "testing"

func TestCalculateRoutes(t *testing.T) {
	lines := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	res := CalculateRoutes(lines)
	if res != 41 {
		t.Errorf("CalculateRoutes() = %v; want 41", res)
	}
}
