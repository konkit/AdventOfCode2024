package day02

import (
	"AdventOfCode2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	arrays := readArrays()

	safe := 0

	for _, arr := range arrays {
		if CheckIfSafeWithDampener(arr) {
			safe++
		}
	}

	fmt.Printf("Safe: %d\n", safe)
}

func readArrays() [][]int {
	file, err := os.Open("daytwo/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)

		numbers := make([]int, len(words))
		for i, word := range words {
			num, _ := strconv.Atoi(word)
			numbers[i] = num
		}

		result = append(result, numbers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func CheckIfSafe(arr []int) bool {
	desc := arr[1]-arr[0] < 0
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		if diff == 0 || utils.Abs(diff) > 3 {
			return false
		}

		if diff < 0 && !desc || diff > 0 && desc {
			return false
		}
	}

	return true
}

func CheckIfSafeWithDampener(arr []int) bool {
	if CheckIfSafe(arr) {
		return true
	}

	for i := 0; i < len(arr); i++ {
		var newArr []int
		for j := 0; j < len(arr); j++ {
			if j == i {
				continue
			}
			newArr = append(newArr, arr[j])
		}

		if CheckIfSafe(newArr) {
			return true
		}
	}

	return false
}
