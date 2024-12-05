package dayfive_test

import (
	"AdventOfCode2024/dayfive"
	"testing"
)

func TestRssult1(t *testing.T) {
	rules := []dayfive.Rule{
		{A: 47, B: 53},
		{A: 97, B: 13},
		{A: 97, B: 61},
		{A: 97, B: 47},
		{A: 75, B: 29},
		{A: 61, B: 13},
		{A: 75, B: 53},
		{A: 29, B: 13},
		{A: 97, B: 29},
		{A: 53, B: 29},
		{A: 61, B: 53},
		{A: 97, B: 53},
		{A: 61, B: 29},
		{A: 47, B: 13},
		{A: 75, B: 47},
		{A: 97, B: 75},
		{A: 47, B: 61},
		{A: 75, B: 61},
		{A: 47, B: 29},
		{A: 75, B: 13},
		{A: 53, B: 13},
	}

	updates := [][]int{
		[]int{75, 47, 61, 53, 29},
		[]int{97, 61, 53, 29, 13},
		[]int{75, 29, 13},
		[]int{75, 97, 47, 61, 53},
		[]int{61, 13, 29},
		[]int{97, 13, 75, 29, 47},
	}

	result := dayfive.Rssult1(rules, updates)

	if result != 143 {
		t.Errorf("Expected 143, got %d", result)
	}
}
