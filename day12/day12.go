package day12

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	input := utils.ReadLines("day12/input.txt")
	res := CalculatePrice(input)
	fmt.Println(res)
}

func CalculatePrice(lines []string) int {
	var plantsMap [][]byte
	plantsMap = generatePlantsMap(lines)

	visitedMap := make(map[int]bool)

	totalPrice := 0

	for y := 0; y < len(plantsMap); y++ {
		for x := 0; x < len(plantsMap[0]); x++ {
			visited := visitedMap[x+y*len(plantsMap[0])]
			if visited {
				continue
			}

			area, perim := traverse(plantsMap, x, y, visitedMap)
			totalPrice += area * perim
		}
	}

	return totalPrice
}

func traverse(plantsMap [][]byte, x int, y int, visitedMap map[int]bool) (int, int) {
	if visitedMap[x+y*len(plantsMap[0])] {
		return 0, 0
	}

	visitedMap[x+y*len(plantsMap[0])] = true

	area := 1
	perimeter := 0

	if y > 0 && plantsMap[y-1][x] == plantsMap[y][x] {
		nArea, nPerim := traverse(plantsMap, x, y-1, visitedMap)
		area += nArea
		perimeter += nPerim
	} else {
		perimeter += 1
	}

	if y < len(plantsMap)-1 && plantsMap[y+1][x] == plantsMap[y][x] {
		nArea, nPerim := traverse(plantsMap, x, y+1, visitedMap)
		area += nArea
		perimeter += nPerim
	} else {
		perimeter += 1
	}

	if x > 0 && plantsMap[y][x-1] == plantsMap[y][x] {
		nArea, nPerim := traverse(plantsMap, x-1, y, visitedMap)
		area += nArea
		perimeter += nPerim
	} else {
		perimeter += 1
	}

	if x < len(plantsMap[0])-1 && plantsMap[y][x+1] == plantsMap[y][x] {
		nArea, nPerim := traverse(plantsMap, x+1, y, visitedMap)
		area += nArea
		perimeter += nPerim
	} else {
		perimeter += 1
	}

	return area, perimeter
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
