package day12

import (
	"AdventOfCode2024/utils"
	"fmt"
)

const UP = "UP"
const DOWN = "DOWN"
const LEFT = "LEFT"
const RIGHT = "RIGHT"

var plantsMap [][]byte

var visitedMap = make(map[int]bool)
var topEdgesMap = make(map[int]bool)
var bottomEdgesMap = make(map[int]bool)
var leftEdgesMap = make(map[int]bool)
var rightEdgesMap = make(map[int]bool)
var visitedEdgesMap = map[string]map[int]bool{
	UP:    topEdgesMap,
	DOWN:  bottomEdgesMap,
	LEFT:  leftEdgesMap,
	RIGHT: rightEdgesMap,
}

var ind = []struct {
	xInc     int
	yInc     int
	dir      string
	edgeDirA string
	edgeDirB string
}{
	{0, -1, UP, LEFT, RIGHT},
	{0, 1, DOWN, LEFT, RIGHT},
	{-1, 0, LEFT, UP, DOWN},
	{1, 0, RIGHT, UP, DOWN},
}

var dirToInd = map[string][]int{
	UP:    {0, -1},
	DOWN:  {0, 1},
	LEFT:  {-1, 0},
	RIGHT: {1, 0},
}

func Run() {
	input := utils.ReadLines("day12/input.txt")
	res := CalculatePrice2(input)
	fmt.Println(res)
}

func CalculatePrice(lines []string) int {
	plantsMap = generatePlantsMap(lines)
	visitedMap = make(map[int]bool)

	totalPrice := 0

	for y := 0; y < len(plantsMap); y++ {
		for x := 0; x < len(plantsMap[0]); x++ {
			visited := visitedMap[x+y*len(plantsMap[0])]
			if visited {
				continue
			}

			area, perim := traverse(x, y)
			totalPrice += area * perim
		}
	}

	return totalPrice
}

func CalculatePrice2(lines []string) int {
	plantsMap = generatePlantsMap(lines)

	totalPrice := 0

	for y := 0; y < len(plantsMap); y++ {
		for x := 0; x < len(plantsMap[0]); x++ {
			visited := visitedMap[x+y*len(plantsMap[0])]
			if visited {
				continue
			}

			area, perim := traverse2(x, y)
			totalPrice += area * perim
		}
	}

	return totalPrice
}

func traverse(x int, y int) (int, int) {
	if isVisited(x, y) {
		return 0, 0
	}

	setVisited(x, y)

	area := 1
	perimeter := 0

	for _, xy := range dirToInd {
		xx := xy[0]
		yy := xy[1]

		if withinBounds(x+xx, y+yy) && plantsMap[y+yy][x+xx] == plantsMap[y][x] {
			nArea, nPerim := traverse(x+xx, y+yy)
			area += nArea
			perimeter += nPerim
		} else {
			perimeter += 1
		}
	}

	return area, perimeter
}

func traverse2(x int, y int) (int, int) {
	if isVisited(x, y) {
		return 0, 0
	}

	setVisited(x, y)

	area := 1
	perimeter := 0

	for _, data := range ind {
		xx := data.xInc
		yy := data.yInc

		if withinBounds(x+xx, y+yy) && plantsMap[y+yy][x+xx] == plantsMap[y][x] {
			nArea, nPerim := traverse2(x+xx, y+yy)
			area += nArea
			perimeter += nPerim
		} else {
			if !isEdgeVisited(x, y, data.dir) {
				perimeter += 1
				setEdgeVisited(x, y, data.dir)

				traverseEdge(x, y, plantsMap[y][x], data.dir, data.edgeDirA)
				traverseEdge(x, y, plantsMap[y][x], data.dir, data.edgeDirB)
			}
		}
	}

	return area, perimeter
}

func traverseEdge(x int, y int, plant byte, edge string, direction string) {
	if plantsMap[y][x] != plant {
		return
	}

	if withinBounds(x+dirToInd[edge][0], y+dirToInd[edge][1]) && plantsMap[y+dirToInd[edge][1]][x+dirToInd[edge][0]] == plant {
		return
	}

	setEdgeVisited(x, y, edge)

	if withinBounds(x+dirToInd[direction][0], y+dirToInd[direction][1]) {
		traverseEdge(x+dirToInd[direction][0], y+dirToInd[direction][1], plant, edge, direction)
	}
}

func setEdgeVisited(x int, y int, edge string) {
	visitedEdgesMap[edge][x+y*len(plantsMap[0])] = true
}

func isEdgeVisited(x, y int, edge string) bool {
	val, ok := visitedEdgesMap[edge][x+y*len(plantsMap[0])]
	if ok {
		return val
	} else {
		return false
	}
}

func setVisited(x int, y int) {
	visitedMap[x+y*len(plantsMap[0])] = true
}

func isVisited(x int, y int) bool {
	return visitedMap[x+y*len(plantsMap[0])]
}

func withinBounds(x int, y int) bool {
	return x >= 0 && x < len(plantsMap[0]) && y >= 0 && y < len(plantsMap)
}

func generatePlantsMap(lines []string) [][]byte {
	var res [][]byte
	for _, line := range lines {
		var lineArr []byte
		for _, char := range line {
			lineArr = append(lineArr, byte(char))
		}
		res = append(res, lineArr)
	}
	return res
}
