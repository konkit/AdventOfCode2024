package day11

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
)

func Run() {
	lines := utils.ReadLines("day11/input.txt")
	result := CalcWithMemo(lines[0], 75)
	fmt.Printf("Number of stones: %d\n", result)
}

var numberToIterationsLeftToCount = make(map[int]map[int]int)

func CalcWithMemo(line string, iterations int) int {
	numbers, err := utils.ParseNumbersOnLine(line)
	if err != nil {
		panic(err)
	}

	count := 0
	for _, initNum := range numbers {
		iterLeft := iterations
		count += countForNumber(initNum, iterLeft)
	}

	return count
}

func countForNumber(e int, left int) int {
	if left == 0 {
		return 1
	}

	i := getNumFromCache(e, left)
	if i != nil {
		return *i
	}

	res := 0
	if e == 0 {
		// If the stone is engraved with the number 0,
		// it is replaced by a stone engraved with the number 1.
		res = countForNumber(1, left-1)
	} else if numOfDigits(e)%2 == 0 {
		// If the stone is engraved with a number that has an even number of digits,
		// it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
		a, b := splitInTwo(e)
		countForA := countForNumber(a, left-1)
		countForB := countForNumber(b, left-1)
		res = countForA + countForB
	} else {
		//If none of the other rules apply, the stone is replaced by a new stone;
		// the old stone's number multiplied by 2024 is engraved on the new stone.
		res = countForNumber(e*2024, left-1)
	}
	setNumToCache(e, left, res)
	return res
}

func getNumFromCache(e int, left int) *int {
	if numberToIterationsLeftToCount[e] == nil {
		return nil
	}
	if val, ok := numberToIterationsLeftToCount[e][left]; ok {
		return &val
	}
	return nil
}

func setNumToCache(e int, left int, count int) {
	if numberToIterationsLeftToCount[e] == nil {
		numberToIterationsLeftToCount[e] = make(map[int]int)
	}
	numberToIterationsLeftToCount[e][left] = count
}

func splitInTwo(i int) (int, int) {
	asStr := strconv.Itoa(i)

	aStr := asStr[0 : len(asStr)/2]
	bStr := asStr[len(asStr)/2:]

	a, err := strconv.Atoi(aStr)
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi(bStr)
	if err != nil {
		fmt.Println(err)
	}
	return a, b
}

func numOfDigits(num int) int {
	if num == 0 {
		return 1
	}
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}
