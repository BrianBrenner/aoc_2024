package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// before is map from a character to a set of characters that need to be seen before
func getMiddle(line string, before map[string]map[string]bool) int {
	chars := strings.Split(line, ",")
	charSet := map[string]bool{}
	for _, char := range chars {
		charSet[char] = true
	}

	seen := map[string]bool{}
	for _, char := range chars {
		reqs := before[char]
		seen[char] = true
		if len(reqs) == 0 {
			continue
		}
		for req := range reqs {
			if charSet[req] && !seen[req] {
				return 0
			}
		}
	}

	middle, _ := strconv.Atoi(chars[len(chars)/2])
	return middle
}

func getRequiredBefore(mappings []string) map[string]map[string]bool {
	before := map[string]map[string]bool{}
	for _, mapping := range mappings {
		nums := strings.Split(mapping, "|")
		if before[nums[1]] == nil {
			before[nums[1]] = map[string]bool{}
		}
		before[nums[1]][nums[0]] = true
	}

	return before
}

func part1() {
	// map from right side to left, ie characters that need to be seen before
	// for each line create a set of all the numbers
	// then create an empty set of visited
	// for each character, check map for values
	// if value is in the line, it must have been seen before therefore it must be in visited

	raw, _ := os.ReadFile("input.txt")
	//raw, _ := os.ReadFile("test.txt")
	data := string(raw)
	parts := strings.Split(data, "\n\n")
	mappings := parts[0]
	lines := parts[1]

	// maps from a number to set of numbers that need to be seen before
	before := getRequiredBefore(strings.Split(mappings, "\n"))
	tot := 0
	for _, line := range strings.Split(lines, "\n") {
		tot += getMiddle(line, before)
	}

	fmt.Println(tot)
}

func part2() {
	raw, _ := os.ReadFile("input.txt")
	//raw, _ := os.ReadFile("test.txt")
	data := string(raw)
	parts := strings.Split(data, "\n\n")
	mappings := parts[0]
	lines := parts[1]

	// maps from a number to set of numbers that need to be seen before
	before := getRequiredBefore(strings.Split(mappings, "\n"))

	broken := []string{}
	for _, line := range strings.Split(lines, "\n") {
		middle := getMiddle(line, before)
		if middle == 0 {
			broken = append(broken, line)
		}
	}

	tot := 0
	for _, broke := range broken {
		arr := strings.Split(broke, ",")
		slices.SortFunc(arr, func(a, b string) int {
			beforeA, ok := before[a]
			// b needs to be before a, ie b < a
			if ok && beforeA[b] {
				return 1
			}
			return -1
		})
		tot += getMiddle(strings.Join(arr, ","), before)
	}

	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
