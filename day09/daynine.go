package day09

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func Run() {
	line := utils.ReadLines("day09/input.txt")[0]
	res := Defragment2(line)
	fmt.Println(res)
}

func Defragment(input string) int {
	blockList := parseList(input)

	start := 0
	end := len(blockList) - 1

	for start < end {
		if blockList[start] == -1 {
			for blockList[end] == -1 {
				end--
			}
			blockList[start] = blockList[end]
			blockList[end] = -1
		}
		start++
	}

	return checkSum(blockList)
}

func parseList(input string) []int {
	var result []int
	fileId := 0
	for i, char := range input {
		num := int(char - '0')
		if i%2 == 0 {
			for i := 0; i < num; i++ {
				result = append(result, fileId)
			}
			fileId++
		} else {
			for i := 0; i < num; i++ {
				result = append(result, -1)
			}
		}
	}
	return result
}

func checkSum(list []int) int {
	sum := 0
	for i, c := range list {
		if c != -1 {
			sum += c * i
		}
	}

	return sum
}

func Defragment2(input string) int {
	blockList := parseList(input)

	end := len(blockList) - 1

	for 0 < end {
		for blockList[end] == -1 {
			end--
		}

		endSize := 0
		for i := end; i >= 0; i-- {
			if blockList[i] != blockList[end] {
				break
			}
			endSize++
		}

		for start := 0; start < end; start++ {
			if blockList[start] != -1 {
				continue
			}

			startEmptySize := 0
			for i := start; i < end; i++ {
				if blockList[i] != -1 {
					break
				}
				startEmptySize++
			}

			if startEmptySize >= endSize {
				for i := 0; i < endSize; i++ {
					blockList[start+i] = blockList[end-i]
					blockList[end-i] = -1
				}
				break
			}
		}

		end -= endSize
	}

	return checkSum(blockList)
}
