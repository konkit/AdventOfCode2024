package dayseven

import (
	"AdventOfCode2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	lines := utils.ReadLines("day07/input.txt")
	result := CheckEquations2(lines)
	fmt.Printf("Result: %d\n", result)
}

func CheckEquations(lines []string) int {
	result := 0
	for _, line := range lines {
		tv, nums := parseLine(line)
		if isCorrect(tv, nums) {
			result += tv
		}
	}

	return result
}

func isCorrect(tv int, nums []int) bool {
	return backTrack(tv, nums[1:], nums[0])
}

func backTrack(tv int, nums []int, result int) bool {
	if len(nums) == 0 {
		return result == tv
	}

	resAdd := backTrack(tv, nums[1:], result+nums[0])
	if resAdd == true {
		return true
	}

	resMul := backTrack(tv, nums[1:], result*nums[0])
	if resMul == true {
		return true
	}

	return false
}

func parseLine(line string) (int, []int) {
	//re := regexp.MustCompile(`^(\d+):( \d+)+`)
	re := regexp.MustCompile(`^(\d+): (\d+(?: \d+)*)`)
	match := re.FindStringSubmatch(line)

	tv, err := strconv.Atoi(match[1])
	if err != nil {
		fmt.Printf("Error converting %s to int\n", match[1])
		return -1, nil
	}

	var nums []int
	numsStr := strings.Split(match[2], " ")
	for _, ns := range numsStr {
		x, err := strconv.Atoi(ns)
		if err != nil {
			fmt.Printf("Error converting %s to int\n", ns)
		}
		nums = append(nums, x)
	}

	return tv, nums
}

func CheckEquations2(lines []string) int {
	result := 0
	for _, line := range lines {
		tv, nums := parseLine(line)
		if isCorrect2(tv, nums) {
			result += tv
		}
	}

	return result
}

func isCorrect2(tv int, nums []int) bool {
	return backTrack2(tv, nums[1:], nums[0])
}

func backTrack2(tv int, nums []int, result int) bool {
	if len(nums) == 0 {
		return result == tv
	}

	resAdd := backTrack2(tv, nums[1:], result+nums[0])
	if resAdd == true {
		return true
	}

	resMul := backTrack2(tv, nums[1:], result*nums[0])
	if resMul == true {
		return true
	}

	resConcat := backTrack2(tv, nums[1:], concat(result, nums[0]))
	if resConcat == true {
		return true
	}

	return false
}

func concat(a int, b int) int {
	str := fmt.Sprintf("%d%d", a, b)
	res, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error converting %s to int\n", str)
	}
	return res
}
