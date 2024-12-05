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

	res := Rssult1(rules, updates)
	fmt.Println(res)
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

type Graph struct {
	adjList  map[int][]int
	inDegree map[int]int
}

func NewGraph() *Graph {
	g := &Graph{
		adjList:  make(map[int][]int),
		inDegree: make(map[int]int),
	}
	//for i := 0; i < n; i++ {
	//	g.adjList[i] = []int{}
	//	g.inDegree[i] = 0
	//}
	return g
}

func (g *Graph) AddEdge(rule Rule) {
	g.adjList[rule.A] = append(g.adjList[rule.A], rule.B)

	val, ok := g.inDegree[rule.A]
	if !ok {
		g.inDegree[rule.A] = 0
	}

	val, ok = g.inDegree[rule.B]
	if ok {
		g.inDegree[rule.B] = val + 1
	} else {
		g.inDegree[rule.B] = 1
	}
}

func (g *Graph) TopologicalSort() []int {
	queue := []int{}
	result := []int{}

	// Find all nodes with in-degree 0
	for node, degree := range g.inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range g.adjList[node] {
			g.inDegree[neighbor]--
			if g.inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If there's a cycle, the graph is not a DAG
	for _, degree := range g.inDegree {
		if degree > 0 {
			return nil
		}
	}

	return result
}

func Rssult1(rules []Rule, updates [][]int) int {
	g := NewGraph()

	for _, rule := range rules {
		g.AddEdge(rule)
	}

	result := g.TopologicalSort()

	sum := 0

	for _, update := range updates {
		correctUpdate := true
		cntVal := getIdOfNode(result, update[0])

		for i := 1; i < len(update) && correctUpdate; i++ {
			x := getIdOfNode(result, update[i])
			if x < cntVal {
				correctUpdate = false
				continue
			}
			cntVal = x
		}

		if correctUpdate {
			sum += getMiddleElement(update)
		}
	}

	fmt.Println(result)
	fmt.Println(sum)

	return sum
}

func getMiddleElement(update []int) int {
	return update[len(update)/2]
}

func getIdOfNode(result []int, node int) int {
	for i := 0; i < len(result); i++ {
		if result[i] == node {
			return i
		}
	}

	return -1
}
