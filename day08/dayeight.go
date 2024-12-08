package dayeight

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	lines := utils.ReadLines("day08/input.txt")
	//result := CalculateAntinodes(lines)
	result := CalculateAntinodes2(lines)
	fmt.Printf("Result: %d\n", result)
}

type Location struct {
	x, y int
}

func CalculateAntinodes(lines []string) int {
	freqToLoc := getFreqToLoc(lines)
	maxX := len(lines)
	maxY := len(lines[0])
	antiNodesMap := make([]bool, maxX*maxY)

	for _, positions := range freqToLoc {
		for i, posI := range positions {
			for j, posJ := range positions {
				if i == j {
					continue
				}

				xDiff := posJ.x - posI.x
				yDiff := posJ.y - posI.y
				an1 := Location{posI.x - xDiff, posI.y - yDiff}
				if withinBounds(maxX, maxY, an1) {
					setAntinode(antiNodesMap, maxY, an1)
				}

				an2 := Location{posJ.x + xDiff, posJ.y + yDiff}
				if withinBounds(maxX, maxY, an2) {
					setAntinode(antiNodesMap, maxY, an2)
				}
			}
		}
	}

	return countAntinodes(antiNodesMap, maxY)
}

func CalculateAntinodes2(lines []string) int {
	freqToLoc := getFreqToLoc(lines)
	maxX := len(lines)
	maxY := len(lines[0])
	antiNodesMap := make([]bool, maxX*maxY)

	for _, positions := range freqToLoc {
		for i, posI := range positions {
			for j, posJ := range positions {
				if i == j {
					continue
				}

				xDiff := posJ.x - posI.x
				yDiff := posJ.y - posI.y

				an1 := Location{posI.x, posI.y}
				for withinBounds(maxX, maxY, an1) {
					setAntinode(antiNodesMap, maxY, an1)
					an1.x = an1.x - xDiff
					an1.y = an1.y - yDiff
				}

				an2 := Location{posJ.x, posJ.y}
				for withinBounds(maxX, maxY, an2) {
					setAntinode(antiNodesMap, maxY, an2)
					an2.x = an2.x + xDiff
					an2.y = an2.y + yDiff
				}
			}
		}
	}

	return countAntinodes(antiNodesMap, maxY)
}

func withinBounds(maxX, maxY int, loc Location) bool {
	if loc.x < 0 || loc.y < 0 || loc.x >= maxX || loc.y >= maxY {
		return false
	}

	return true
}

func getFreqToLoc(lines []string) map[byte][]Location {
	m := map[byte][]Location{}

	for x, line := range lines {
		for y := 0; y < len(line); y++ {
			char := line[y]

			if char != '.' {
				l := Location{x, y}
				m[char] = append(m[char], l)
			}
		}
	}

	return m
}

func setAntinode(isAntinode []bool, lineSize int, an1 Location) {
	isAntinode[an1.x*lineSize+an1.y] = true
}

func countAntinodes(antinode []bool, lineSize int) int {
	sum := 0
	for i, an := range antinode {
		if an == true {
			fmt.Printf("Antinode in %dx%d\n", i/lineSize, i%lineSize)
			sum++
		}
	}
	return sum
}
