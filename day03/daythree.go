package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	content, err := os.ReadFile("daythree/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string
	fileContent := string(content)

	result := DoMultiplyPartTwo(fileContent)
	fmt.Printf("Result: %d\n", result)
}

func DoMultiply(input string) int {
	m := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	results := m.FindAllString(input, -1)

	sum := 0

	for _, str := range results {
		a, b := parseMultiply(str, sum)
		sum += a * b
	}

	fmt.Println(results)

	return sum
}

func DoMultiplyPartTwo(input string) int {
	m := regexp.MustCompile(`(mul\(([0-9]+),([0-9]+)\))|(do\(\))|(don't\(\))`)

	results := m.FindAllString(input, -1)

	sum := 0
	enabled := true

	for _, str := range results {
		if strings.HasPrefix(str, "don't") {
			enabled = false
		} else if strings.HasPrefix(str, "do") {
			enabled = true
		} else {
			if enabled {
				a, b := parseMultiply(str, sum)
				sum += a * b
			}
		}
	}

	fmt.Println(results)

	return sum
}

func parseMultiply(str string, sum int) (int, int) {
	mm := regexp.MustCompile(`mul\((\d{1,4}),(\d{1,4})\)`)

	r := mm.MatchString(str)
	fmt.Println(r)

	res := mm.FindAllStringSubmatch(str, -1)
	a, err := strconv.Atoi(res[0][1])
	b, err := strconv.Atoi(res[0][2])

	if err != nil {
		fmt.Println(err)
	}

	return a, b
}
