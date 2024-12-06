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

func TestCalculateObstacle(t *testing.T) {
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
	res := CalculateObstacle(lines, 100)
	if res != 6 {
		t.Errorf("CalculateObstacle() = %v; want 6", res)
	}
}
