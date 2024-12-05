package dayfive

import (
	"AdventOfCode2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Run() {
	rules, updates, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	res := Rssult2(rules, updates)
	fmt.Println(res)
}

func Rssult1(rules []Rule, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		correctUpdate := checkUpdate(rules, update)

		if correctUpdate {
			sum += getMiddleElement(update)
		}
	}

	return sum
}

func Rssult2(rules []Rule, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		isCorrect := checkUpdate(rules, update)

		if isCorrect {
			//sum += getMiddleElement(update)
		} else {
			makeUpdateCorrect(update, rules)
			sum += getMiddleElement(update)
		}
	}

	return sum
}

func makeUpdateCorrect(update []int, rules []Rule) {
	changed := true
	for changed == true {
		changed = false

		for _, rule := range rules {
			aPos := findElement(update, rule.A)
			if aPos == -1 {
				continue
			}

			bPos := findElement(update, rule.B)
			if bPos == -1 {
				continue
			}

			if aPos > bPos {
				swap(update, aPos, bPos)
				changed = true
			}
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func findElement(update []int, val int) int {
	for i := 0; i < len(update); i++ {
		if update[i] == val {
			return i
		}
	}
	return -1
}

func checkUpdate(rules []Rule, update []int) bool {
	entryToPosition := map[int]int{}

	for i := 0; i < len(update); i++ {
		entryToPosition[update[i]] = i
	}

	for _, rule := range rules {
		aPos, ok := entryToPosition[rule.A]
		if !ok {
			continue
		}

		bPos, ok := entryToPosition[rule.B]
		if !ok {
			continue
		}

		if aPos > bPos {
			return false
		}
	}

	return true
}

func readInput() ([]Rule, [][]int, error) {
	lines := utils.ReadLines("dayfive/input.txt")

	var rules []Rule
	var updates [][]int
	readingRules := true

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			readingRules = false
			continue
		}

		if readingRules {
			var a, b int
			_, err := fmt.Sscanf(line, "%d|%d", &a, &b)
			if err != nil {
				fmt.Printf("Error parsing line %s: %s\n", line, err)
				return nil, nil, err
			}

			rules = append(rules, Rule{a, b})
		} else {
			var update []int
			res := strings.Split(line, ",")
			for _, a := range res {
				num, err := strconv.Atoi(a)
				if err != nil {
					fmt.Printf("Error parsing line %s: %s\n", line, err)
					return nil, nil, err
				}
				update = append(update, num)
			}

			updates = append(updates, update)
		}
	}
	return rules, updates, nil
}

type Rule struct {
	A int
	B int
}

func getMiddleElement(update []int) int {
	return update[len(update)/2]
}
