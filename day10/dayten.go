package dayten

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	lines := utils.ReadLines("day10/input.txt")
	res := Calculate(lines)
	fmt.Println(res)
}

type Pos struct {
	x, y int
}

func Calculate(lines []string) int {
	topMap := constructMap(lines)

	scoresSum := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if topMap[y][x] != 0 {
				continue
			}

			visited := map[int]bool{}

			//score := calculateScore(topMap, visited, x, y)
			score := calculateRating(topMap, visited, x, y)
			scoresSum += score
		}
	}

	return scoresSum
}

func calculateScore(topMap [][]int, visited map[int]bool, x int, y int) int {
	if topMap[y][x] == 9 {
		res, ok := visited[y*len(topMap[x])+x]
		if ok && res == true {
			return 0
		} else {
			visited[y*len(topMap[x])+x] = true
			return 1
		}
	}

	score := 0
	if y > 0 && topMap[y-1][x] == topMap[y][x]+1 {
		score += calculateScore(topMap, visited, x, y-1)
	}
	if x > 0 && topMap[y][x-1] == topMap[y][x]+1 {
		score += calculateScore(topMap, visited, x-1, y)
	}
	if y < len(topMap)-1 && topMap[y+1][x] == topMap[y][x]+1 {
		score += calculateScore(topMap, visited, x, y+1)
	}
	if x < len(topMap[0])-1 && topMap[y][x+1] == topMap[y][x]+1 {
		score += calculateScore(topMap, visited, x+1, y)
	}
	return score
}

func calculateRating(topMap [][]int, visited map[int]bool, x int, y int) int {
	if topMap[y][x] == 9 {
		return 1
		//res, ok := visited[y*len(topMap[x])+x]
		//if ok && res == true {
		//	return 0
		//} else {
		//	visited[y*len(topMap[x])+x] = true
		//	return 1
		//}
	}

	score := 0
	if y > 0 && topMap[y-1][x] == topMap[y][x]+1 {
		score += calculateRating(topMap, visited, x, y-1)
	}
	if x > 0 && topMap[y][x-1] == topMap[y][x]+1 {
		score += calculateRating(topMap, visited, x-1, y)
	}
	if y < len(topMap)-1 && topMap[y+1][x] == topMap[y][x]+1 {
		score += calculateRating(topMap, visited, x, y+1)
	}
	if x < len(topMap[0])-1 && topMap[y][x+1] == topMap[y][x]+1 {
		score += calculateRating(topMap, visited, x+1, y)
	}
	return score
}

func constructMap(lines []string) [][]int {
	topMap := [][]int{}

	for _, line := range lines {
		arr := make([]int, len(line))
		for x, c := range line {
			arr[x] = int(c - '0')
		}
		topMap = append(topMap, arr)
	}

	return topMap
}
