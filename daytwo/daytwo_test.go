package daytwo_test

import (
	"AdventOfCode2024/daytwo"
	"testing"
)

func TestCheckIfSafe(t *testing.T) {
	testcase := []struct {
		name string
		arr  []int
		safe bool
	}{
		{
			name: "testcase1",
			arr:  []int{7, 6, 4, 2, 1},
			safe: true,
		},
		{
			name: "testcase2",
			arr:  []int{1, 2, 7, 8, 9},
			safe: false,
		},
		{
			name: "testcase3",
			arr:  []int{9, 7, 6, 2, 1},
			safe: false,
		},
		{
			name: "testcase4",
			arr:  []int{1, 3, 2, 4, 5},
			safe: false,
		},
		{
			name: "testcase5",
			arr:  []int{8, 6, 4, 4, 1},
			safe: false,
		},
		{
			name: "testcase6",
			arr:  []int{1, 3, 6, 7, 9},
			safe: true,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			res := daytwo.CheckIfSafe(tc.arr)
			if res != tc.safe {
				t.Errorf("CheckIfSafe was incorrect, got: %v, want: %v.", tc.arr, tc.safe)
			}
		})
	}
}

func TestCheckIfSafeWithDampener(t *testing.T) {
	testcase := []struct {
		name string
		arr  []int
		safe bool
	}{
		{
			name: "testcase1",
			arr:  []int{7, 6, 4, 2, 1},
			safe: true,
		},
		{
			name: "testcase2",
			arr:  []int{1, 2, 7, 8, 9},
			safe: false,
		},
		{
			name: "testcase3",
			arr:  []int{9, 7, 6, 2, 1},
			safe: false,
		},
		{
			name: "testcase4",
			arr:  []int{1, 3, 2, 4, 5},
			safe: true,
		},
		{
			name: "testcase5",
			arr:  []int{8, 6, 4, 4, 1},
			safe: true,
		},
		{
			name: "testcase6",
			arr:  []int{1, 3, 6, 7, 9},
			safe: true,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			res := daytwo.CheckIfSafeWithDampener(tc.arr)
			if res != tc.safe {
				t.Errorf("CheckIfSafeWithDampener was incorrect, got: %v, want: %v.", tc.arr, tc.safe)
			}
		})
	}
}
