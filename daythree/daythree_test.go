package daythree_test

import (
	"AdventOfCode2024/daythree"
	"testing"
)

func TestDoMultiply(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expectedRes := 161

	res := daythree.DoMultiply(input)

	if res != expectedRes {
		t.Errorf("got %d, want %d", res, expectedRes)
	}
}

func TestDoMultiplyPartTwo(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))\n)"
	expectedRes := 48

	res := daythree.DoMultiplyPartTwo(input)

	if res != expectedRes {
		t.Errorf("got %d, want %d", res, expectedRes)
	}
}
