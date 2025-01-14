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

func cloneMap(src map[string]bool) map[string]bool {
	dst := make(map[string]bool, len(src))
	for key, val := range src {
		dst[key] = val
	}
	return dst
}

func dfs(cur string, root string, graph map[string]map[string]bool, path map[string]bool) int {
	// found loop with root, return len of loop
	if cur == root && len(path) > 0 {
		return len(path) - 1
	}
	// visited
	_, ok := path[cur]
	if ok && len(path) > 0 {
		return 0
	}

	maxVal := 0
	for next, _ := range graph[cur] {
		newPath := cloneMap(path)
		newPath[cur] = true

		maxVal = max(maxVal, dfs(next, root, graph, newPath))
	}

	return maxVal
}

// [a, b,c,d] would return [[a,b],[a,c],[a,d],[b,c],[b,d],[c,d]]
//func getAllCombos(vals []string) [][]string {
//	var combos [][]string
//	for i := 0; i < len(vals); i++ {
//		for j := i + 1; j < len(vals); j++ {
//			combos = append(combos, []string{vals[i], vals[j]})
//		}
//	}
//	return combos
//}

// chatgpt
func getAllCombos(vals []string) [][]string {
	var result [][]string

	// current will hold the subset we're building at each recursion step
	var current []string

	// backtrack is a recursive function that tries including and excluding
	// each element at index i.
	var backtrack func(i int)
	backtrack = func(i int) {
		// If we've gone past the last element, save the current subset.
		if i == len(vals) {
			// Make a copy of current so we don't overwrite it in future calls
			subsetCopy := make([]string, len(current))
			copy(subsetCopy, current)
			result = append(result, subsetCopy)
			return
		}

		// 1) Exclude vals[i] and recurse
		backtrack(i + 1)

		// 2) Include vals[i] and recurse
		current = append(current, vals[i])
		backtrack(i + 1)

		// Remove the last included element to backtrack properly
		current = current[:len(current)-1]
	}

	// Kick off recursion from the 0th index
	backtrack(0)

	return result
}

// ex: [ka, co, de, ta].
// check if ka has co, de, and ta.
// check if co, has de, ta (co is guarenteed to have ka from previous step)
// check if de has ta
func isLan(combo []string, graph map[string]map[string]bool) bool {
	for i, comp := range combo[:len(combo)-1] {
		have := graph[comp]
		for _, comp2 := range combo[i+1:] {
			_, ok := have[comp2]
			if !ok {
				return false
			}
		}
	}

	return true
}

func findBiggestLan(node string, graph map[string]map[string]bool) []string {
	set := []string{node}
	for edge := range graph[node] {
		set = append(set, edge)
	}
	raw := getAllCombos(set)
	var combos [][]string
	for _, subset := range raw {
		if len(subset) >= 2 {
			combos = append(combos, subset)
		}
	}

	maxCombo := []string{}
	for _, combo := range combos {
		if isLan(combo, graph) {
			if len(combo) > len(maxCombo) {
				//out  = len(combo)
				maxCombo = combo
			}
		}
	}

	return maxCombo

}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")

	graph := createGraph(data)
	biggest := []string{}
	for node := range graph {
		lan := findBiggestLan(node, graph)
		//out = max(out, findBiggestLan(node, graph))
		if len(lan) > len(biggest) {
			biggest = lan
		}
	}
	slices.Sort(biggest)
	fmt.Println(strings.Join(biggest, ","))
}

func main() {
	//part1()
	part2()
}
