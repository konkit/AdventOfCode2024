package dayfour

import "testing"

func TestCount(t *testing.T) {
	lines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	res := Count(lines)
	if res != 18 {
		t.Errorf("want 18, got %d", res)
	}
}

func TestCount2(t *testing.T) {
	lines := []string{
		".M.S......",
		"..A..MSMS.",
		".M.S.MAA..",
		"..A.ASMSM.",
		".M.S.M....",
		"..........",
		"S.S.S.S.S.",
		".A.A.A.A..",
		"M.M.M.M.M.",
		"..........",
	}

	res := Count2(lines)
	if res != 18 {
		t.Errorf("want 18, got %d", res)
	}
}
