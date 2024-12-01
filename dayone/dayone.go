package dayone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

import "golang.org/x/exp/constraints"

func Run() int {
	arr1, arr2 := readArrays()

	//return partOne(arr1, arr2)
	return partTwo(arr1, arr2)
}

func partOne(arr1 []int, arr2 []int) int {
	sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })

	size := 0
	for i := 0; i < len(arr1); i++ {
		size += Abs(arr1[i] - arr2[i])
	}

	return size
}

func partTwo(arr1 []int, arr2 []int) int {
	simScoreMap := make(map[int]int)

	for _, i := range arr2 {
		simScoreMap[i]++
	}

	result := 0

	for _, i := range arr1 {
		num, ok := simScoreMap[i]
		if ok {
			result += i * num
		}
	}

	return result
}

func readArrays() ([]int, []int) {
	file, err := os.Open("dayone/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr1, arr2 []int

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		var a, b int
		_, err = fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			log.Fatal(err)
		}

		arr1 = append(arr1, a)
		arr2 = append(arr2, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return arr1, arr2
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
