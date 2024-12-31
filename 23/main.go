package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func addConnected(computers []string, connected *map[string]bool) {
	// deduplicates by sorting
	slices.Sort(computers)
	key := strings.Join(computers, ",")
	(*connected)[key] = true
}

func findConnected(node string, graph map[string]map[string]bool, connected *map[string]bool) {
	for edge := range graph[node] {
		for subEdge := range graph[edge] {
			if subEdge == node {
				continue
			}
			_, ok := graph[subEdge][node]
			if ok {
				conn := []string{node, edge, subEdge}
				addConnected(conn, connected)
			}
		}
	}
}

func createGraph(data []string) map[string]map[string]bool {
	graph := map[string]map[string]bool{}
	for _, line := range data {
		nodes := strings.Split(line, "-")
		_, ok := graph[nodes[0]]
		if !ok {
			graph[nodes[0]] = map[string]bool{}
		}
		graph[nodes[0]][nodes[1]] = true

		_, ok = graph[nodes[1]]
		if !ok {
			graph[nodes[1]] = map[string]bool{}
		}
		graph[nodes[1]][nodes[0]] = true
	}

	return graph
}

func getTCount(connected map[string]bool) int {
	count := 0
	for conn := range connected {
		comps := strings.Split(conn, ",")
		for _, comp := range comps {
			if strings.HasPrefix(comp, "t") {
				count += 1
				break
			}
		}
	}

	return count
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")

	graph := createGraph(data)
	connected := map[string]bool{}
	for node := range graph {
		findConnected(node, graph, &connected)
	}

	fmt.Println(getTCount(connected))
}

func main() {
	part1()
}
