package day09

import "testing"

func TestDefragment(t *testing.T) {
	input := "2333133121414131402"

	res := Defragment(input)

	if res != 1928 {
		t.Errorf("Defragment result was incorrect, got: %d, want: %d.", res, 1928)
	}
}

func TestDefragment2(t *testing.T) {
	input := "2333133121414131402"

	res := Defragment2(input)

	if res != 2858 {
		t.Errorf("Defragment result was incorrect, got: %d, want: %d.", res, 2858)
	}
}
