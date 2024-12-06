package day06

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	lines := utils.ReadLines("day06/input.txt")
	res := CalculateRoutes(lines)
	fmt.Println(res)
}

func CalculateRoutes(lines []string) int {
	var labmap [][]byte
	for _, line := range lines {
		var newArr []byte
		for _, char := range line {
			newArr = append(newArr, byte(char))
		}
		labmap = append(labmap, newArr)
	}

	visited := map[int]bool{}
	visitedCount := 0

	dir := "up"

	x := 0
	y := 0

	for i := 0; i < len(labmap); i++ {
		for j := 0; j < len(labmap[i]); j++ {
			if labmap[i][j] == '^' {
				x = i
				y = j
			}
		}
	}

	for x > 0 && y > 0 && x < len(labmap) && y < len(labmap[0]) {
		fmt.Printf("Visiting %d %d\n", y, x)
		if visited[x*len(labmap[0])+y] == false {
			visited[x*len(labmap[0])+y] = true
			visitedCount++
		}
		// If there is something directly in front of you, turn right 90 degrees.
		// Otherwise, take a step forward.
		newX, newY, newDir := getNextPlace(labmap, x, y, dir)
		x = newX
		y = newY
		dir = newDir
	}
	return visitedCount
}

func getNextPlace(labmap [][]byte, x int, y int, dir string) (int, int, string) {
	switch dir {
	case "up":
		if x > 0 && labmap[x-1][y] == '#' {
			dir = "right"
			return x, y + 1, dir
		} else {
			return x - 1, y, dir
		}
		break

	case "right":
		if y < len(labmap[0])-1 && labmap[x][y+1] == '#' {
			dir = "down"
			return x + 1, y, dir
		} else {
			return x, y + 1, dir
		}
		break

	case "down":
		if x < len(labmap)-1 && labmap[x+1][y] == '#' {
			dir = "left"
			return x, y - 1, dir
		} else {
			return x + 1, y, dir
		}
		break
	case "left":
		if y > 0 && labmap[x][y-1] == '#' {
			dir = "up"
			return x - 1, y, dir
		} else {
			return x, y - 1, dir
		}
		break
	}

	return 0, 0, dir
}
