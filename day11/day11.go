package day11

import (
	"AdventOfCode2024/utils"
	"container/list"
	"fmt"
	"strconv"
)

func Run() {
	lines := utils.ReadLines("day11/input.txt")
	result := Calculate(lines[0])
	fmt.Printf("Number of stones: %d\n", result)
}

func Calculate(line string) int {
	numbers, err := utils.ParseNumbersOnLine(line)
	if err != nil {
		panic(err)
	}

	stones := list.New()
	for _, number := range numbers {
		stones.PushBack(number)
	}

	iterations := 25
	for i := 0; i < iterations; i++ {
		for e := stones.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == 0 {
				// If the stone is engraved with the number 0,
				// it is replaced by a stone engraved with the number 1.
				e.Value = 1
			} else if numOfDigits(e.Value.(int))%2 == 0 {
				// If the stone is engraved with a number that has an even number of digits,
				// it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
				a, b := splitInTwo(e.Value.(int))
				e.Value = a
				stones.InsertAfter(b, e)
				e = e.Next()
			} else {
				//If none of the other rules apply, the stone is replaced by a new stone;
				// the old stone's number multiplied by 2024 is engraved on the new stone.
				e.Value = e.Value.(int) * 2024
			}

			//listAsStr := utils.ListAsString(stones)
			//fmt.Println(listAsStr)
		}

		//listAsStr := utils.ListAsString(stones)
		//fmt.Println(listAsStr)
	}

	return stones.Len()
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
