package day06

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	lines := utils.ReadLines("day06/input.txt")
	//res := CalculateRoutes(lines)
	res := CalculateObstacle(lines, 100000)
	fmt.Println(res)
}

type Position struct {
	x, y int
}

func CalculateRoutes(lines []string) int {
	labmap := parseLabMap(lines)

	visited := map[int]bool{}
	visitedCount := 0

	dir := "up"
	x, y := getStartingPoint(labmap)

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

func CalculateObstacle(lines []string, threshold int) int {
	labmap := parseLabMap(lines)
	initX, initY := getStartingPoint(labmap)

	candidates := getVisitedPositionsList(labmap)

	possibleObstaclesCount := 0
	for _, candidate := range candidates {
		//for i := 0; i < len(labmap); i++ {
		//	for j := 0; j < len(labmap[i]); j++ {
		obsI, obsJ := candidate.x, candidate.y
		//obsI, obsJ := i, j
		if obsI == initX && obsJ == initY {
			continue
		}
		if labmap[obsI][obsJ] == '#' {
			continue
		}

		// Set new obstacle
		labmap[obsI][obsJ] = '#'

		// Check
		x, y := initX, initY
		dir := "up"
		steps := 0

		for steps < threshold && x >= 0 && y >= 0 && x < len(labmap) && y < len(labmap[0]) {
			newX, newY, newDir := getNextPlace(labmap, x, y, dir)
			x = newX
			y = newY
			dir = newDir

			steps++
		}

		if steps == threshold {
			//fmt.Printf("Loop obstacle: %d %d\n", x, y)
			possibleObstaclesCount++
		}

		// Cleanup
		labmap[obsI][obsJ] = '.'
		//}
	}

	return possibleObstaclesCount
}

func getVisitedPositionsList(labmap [][]byte) []Position {
	var positions []Position
	visited := map[int]bool{}

	dir := "up"
	x, y := getStartingPoint(labmap)

	for x > 0 && y > 0 && x < len(labmap) && y < len(labmap[0]) {
		//fmt.Printf("Visiting %d %d\n", y, x)
		if visited[x*len(labmap[0])+y] == false {
			visited[x*len(labmap[0])+y] = true
			positions = append(positions, Position{x, y})
		}
		// If there is something directly in front of you, turn right 90 degrees.
		// Otherwise, take a step forward.
		newX, newY, newDir := getNextPlace(labmap, x, y, dir)
		x = newX
		y = newY
		dir = newDir
	}

	return positions
}

func parseLabMap(lines []string) [][]byte {
	var labmap [][]byte
	for _, line := range lines {
		var newArr []byte
		for _, char := range line {
			newArr = append(newArr, byte(char))
		}
		labmap = append(labmap, newArr)
	}
	return labmap
}

func getStartingPoint(labmap [][]byte) (int, int) {
	for i := 0; i < len(labmap); i++ {
		for j := 0; j < len(labmap[i]); j++ {
			if labmap[i][j] == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func getNextPlace(labmap [][]byte, x int, y int, dir string) (int, int, string) {
	switch dir {
	case "up":
		if x > 0 && labmap[x-1][y] == '#' {
			dir = "right"
			return x, y, dir
		} else {
			return x - 1, y, dir
		}
		break

	case "right":
		if y < len(labmap[0])-1 && labmap[x][y+1] == '#' {
			dir = "down"
			return x, y, dir
		} else {
			return x, y + 1, dir
		}
		break

	case "down":
		if x < len(labmap)-1 && labmap[x+1][y] == '#' {
			dir = "left"
			return x, y, dir
		} else {
			return x + 1, y, dir
		}
		break
	case "left":
		if y > 0 && labmap[x][y-1] == '#' {
			dir = "up"
			return x, y, dir
		} else {
			return x, y - 1, dir
		}
		break
	}

	return 0, 0, dir
}
