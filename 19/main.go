package main

import (
	"fmt"
	"os"
	"strings"
)

func dfs(towel string, patterns map[string]bool, start int, end int, cache *map[int]bool) bool {
	cacheValue, ok := (*cache)[start]
	if ok {
		return cacheValue
	}

	toCheck := towel[start:end]
	_, ok = patterns[toCheck]
	if !ok {
		return false
	}
	// were able to reach the end with valid towels
	if end == len(towel) {
		return true
	}
	// since it was possible to get to this point, then do DFS for the combinations for the rest of the towel
	for i := end + 1; i <= len(towel); i++ {
		valid := dfs(towel, patterns, end, i, cache)
		if valid == true {
			return valid
		}
	}

	// this means it wasn't possible to find a valid set of towels if you were to start at "end"
	(*cache)[end] = false
	return false
}

func isValid(towel string, patterns map[string]bool) int {
	// map from an index to if is possible to reach the end
	cache := map[int]bool{}
	for i := 1; i <= len(towel); i++ {
		// initialize dfs with different starting patterns. ie brwrr, do dfs for b, br, brw, brwr, and brwrr
		valid := dfs(towel, patterns, 0, i, &cache)
		if valid == true {
			return 1
		}
	}

	return 0
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	rawPatterns := strings.Split(data, "\n\n")[0]
	patterns := map[string]bool{}
	for _, p := range strings.Split(rawPatterns, ", ") {
		patterns[p] = true
	}

	towels := strings.Split(data, "\n\n")[1]
	tot := 0
	for _, towel := range strings.Split(towels, "\n") {
		tot += isValid(towel, patterns)
	}
	fmt.Println(tot)
}

// use DP. start from the back and work to the front, brwrr as an example:
// at index 4, there is one way to reach the end, r
// at index 3, you could do rr (doesn't exist), or r. If r works (it does) then the number of ways to reach the end is 1 + counts[4]
// at index 2, you could do wrr (doesn't exist), or wr which gives 1 + counts[4], or w (doesn't exist) = 1 + counts[3]
// ...
func getCounts(towel string, patterns map[string]bool) int {
	// map from index to the number of ways to reach the end. use DP
	counts := map[int]int{len(towel): 1}
	for i := len(towel) - 1; i >= 0; i-- {
		for j := i + 1; j <= len(towel); j++ {
			toCheck := towel[i:j]
			_, ok := patterns[toCheck]
			if ok {
				counts[i] = counts[i] + counts[j]
			}
		}
	}

	return counts[0]
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	rawPatterns := strings.Split(data, "\n\n")[0]
	patterns := map[string]bool{}
	for _, p := range strings.Split(rawPatterns, ", ") {
		patterns[p] = true
	}

	towels := strings.Split(data, "\n\n")[1]
	tot := 0
	for _, towel := range strings.Split(towels, "\n") {
		tot += getCounts(towel, patterns)
	}
	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
