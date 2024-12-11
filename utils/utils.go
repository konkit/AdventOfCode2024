package utils

import (
	"bufio"
	"container/list"
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
	"os"
	"strconv"
	"strings"
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

func ParseNumbersOnLine(line string) ([]int, error) {
	var res []int
	numbers := strings.Split(line, " ")
	for _, a := range numbers {
		num, err := strconv.Atoi(a)
		if err != nil {
			fmt.Printf("Error parsing line %s: %s\n", line, err)
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}

func ListAsString(l *list.List) string {
	str := ""
	str += "List: "
	for e := l.Front(); e != nil; e = e.Next() {
		str += fmt.Sprintf("%v ", e.Value)
	}
	return str
}
