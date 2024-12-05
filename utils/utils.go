package utils

import (
	"bufio"
	"golang.org/x/exp/constraints"
	"log"
	"os"
)

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
