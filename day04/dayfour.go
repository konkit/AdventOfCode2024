package day04

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	lines := utils.ReadLines("dayfour/input.txt")
	res := Count2(lines)
	fmt.Println(res)
}

func Count(lines []string) int {
	input := parseToArray(lines)
	count := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if CheckXMASVertically(input, i, j, 0) {
				count++
			}
			if CheckXMASHorizontally(input, i, j, 0) {
				count++
			}
			if CheckXMASDiagonally(input, i, j, 0) {
				count++
			}
			if CheckXMASVertically2(input, i, j, 0) {
				count++
			}
			if CheckXMASHorizontally2(input, i, j, 0) {
				count++
			}
			if CheckXMASDiagonally2(input, i, j, 0) {
				count++
			}
			if CheckXMASDiagonally3(input, i, j, 0) {
				count++
			}
			if CheckXMASDiagonally4(input, i, j, 0) {
				count++
			}
		}
	}

	return count
}

func CheckXMASVertically(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASVertically(input, i, j+1, letter+1)
}

func CheckXMASVertically2(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASVertically2(input, i, j-1, letter+1)
}

func CheckXMASHorizontally(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASHorizontally(input, i+1, j, letter+1)
}

func CheckXMASHorizontally2(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASHorizontally2(input, i-1, j, letter+1)
}

func CheckXMASDiagonally(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASDiagonally(input, i+1, j+1, letter+1)
}

func CheckXMASDiagonally2(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASDiagonally2(input, i-1, j-1, letter+1)
}
func CheckXMASDiagonally3(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASDiagonally3(input, i+1, j-1, letter+1)
}
func CheckXMASDiagonally4(input [][]byte, i int, j int, letter int) bool {
	if letter == 4 {
		return true
	}
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != GetXMASLetter(letter) {
		return false
	}
	return CheckXMASDiagonally4(input, i-1, j+1, letter+1)
}

func checkBounds(input [][]byte, i int, j int) bool {
	if i < 0 {
		return false
	}
	if j < 0 {
		return false
	}
	if i >= len(input) {
		return false
	}
	if j >= len(input[0]) {
		return false
	}
	return true
}

func GetXMASLetter(index int) byte {
	const XMAS = "XMAS"
	return XMAS[index]
}

func parseToArray(lines []string) [][]byte {
	var input [][]byte
	for _, line := range lines {
		var lineArr []byte
		for _, b := range []byte(line) {
			lineArr = append(lineArr, b)
		}
		input = append(input, lineArr)
	}
	return input
}

func Count2(lines []string) int {
	input := parseToArray(lines)
	count := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != 'A' {
				continue
			}

			if CheckTop(input, i, j) {
				count++
			}
			if CheckBottom(input, i, j) {
				count++
			}
			if CheckLeft(input, i, j) {
				count++
			}
			if CheckRight(input, i, j) {
				count++
			}
		}
	}

	return count
}

// M M
//
//	A
//
// S S
func CheckTop(input [][]byte, i int, j int) bool {
	if !checkBoundsAndVal(input, i-1, j-1, 'M') {
		return false
	}
	if !checkBoundsAndVal(input, i-1, j+1, 'M') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j-1, 'S') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j+1, 'S') {
		return false
	}
	return true
}

// M S
//
//	A
//
// M S
func CheckLeft(input [][]byte, i int, j int) bool {
	if !checkBoundsAndVal(input, i-1, j-1, 'M') {
		return false
	}
	if !checkBoundsAndVal(input, i-1, j+1, 'S') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j-1, 'M') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j+1, 'S') {
		return false
	}
	return true
}

// S S
//
//	A
//
// M M
func CheckBottom(input [][]byte, i int, j int) bool {
	if !checkBoundsAndVal(input, i-1, j-1, 'S') {
		return false
	}
	if !checkBoundsAndVal(input, i-1, j+1, 'S') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j-1, 'M') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j+1, 'M') {
		return false
	}
	return true
}

// S M
//
//	A
//
// S M
func CheckRight(input [][]byte, i int, j int) bool {
	if !checkBoundsAndVal(input, i-1, j-1, 'S') {
		return false
	}
	if !checkBoundsAndVal(input, i-1, j+1, 'M') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j-1, 'S') {
		return false
	}
	if !checkBoundsAndVal(input, i+1, j+1, 'M') {
		return false
	}
	return true
}

func checkBoundsAndVal(input [][]byte, i int, j int, val byte) bool {
	if !checkBounds(input, i, j) {
		return false
	}
	if input[i][j] != val {
		return false
	}
	return true
}
