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
	if end == len(towel) {
		return true
	}
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
	// map from start to if is valid
	cache := map[int]bool{}
	for i := 1; i <= len(towel); i++ {
		valid := dfs(towel, patterns, 0, i, &cache)
		if valid == true {
			return 1
		}
	}

	return 0
}

func dfs2(towel string, patterns map[string]bool, start int, end int, cache *map[int]bool, count *int) bool {
	cacheValue, ok := (*cache)[start]
	if ok {
		// TODO: check if true maybe?
		return cacheValue
	}

	toCheck := towel[start:end]
	_, ok = patterns[toCheck]
	if !ok {
		return false
	}
	if end == len(towel) {
		*count += 1
		return true
	}
	found := false
	for i := end + 1; i <= len(towel); i++ {
		val := dfs2(towel, patterns, end, i, cache, count)
		if val {
			found = true
		}
	}

	if !found {
		(*cache)[end] = found
	}
	return found
}

func isValid2(towel string, patterns map[string]bool) int {
	// map from start to if is valid
	cache := map[int]bool{}
	count := 0
	for i := 1; i <= len(towel); i++ {
		dfs2(towel, patterns, 0, i, &cache, &count)
	}

	return count
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

func part2() {
	raw, _ := os.ReadFile("test.txt")
	//raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	rawPatterns := strings.Split(data, "\n\n")[0]
	patterns := map[string]bool{}
	for _, p := range strings.Split(rawPatterns, ", ") {
		patterns[p] = true
	}

	towels := strings.Split(data, "\n\n")[1]
	tot := 0
	for _, towel := range strings.Split(towels, "\n") {
		tot += isValid2(towel, patterns)
		//fmt.Println(towel)
		//fmt.Println(isValid2(towel, patterns))
		//fmt.Println("***")
	}
	fmt.Println(tot)
}

func main() {
	//part1()
	part2()
}
