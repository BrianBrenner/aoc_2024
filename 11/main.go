package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(stones []int) []int {
	newStones := []int{}
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}
		strNum := strconv.Itoa(stone)
		if len(strNum)%2 == 0 {
			l, _ := strconv.Atoi(strNum[:len(strNum)/2])
			r, _ := strconv.Atoi(strNum[len(strNum)/2:])
			newStones = append(newStones, l, r)
			continue
		}
		newStones = append(newStones, stone*2024)
	}

	return newStones
}

var stoneCache = map[int][]int{
	0: []int{2, 0, 2, 4},
	1: []int{2, 0, 2, 4},
	2: []int{4, 0, 4, 8},
	3: []int{6, 0, 7, 2},
	4: []int{8, 0, 9, 6},
	5: []int{2, 0, 4, 8, 2, 8, 8, 0},
	6: []int{2, 4, 5, 7, 9, 4, 5, 6},
	7: []int{2, 8, 6, 7, 6, 0, 3, 2},
	8: []int{3, 2, 7, 7, 2, 6, 16192},
	9: []int{3, 6, 8, 6, 9, 1, 8, 4},
}

var blinkCache = map[int]int{
	0: 4,
	1: 3,
	2: 3,
	3: 3,
	4: 3,
	5: 5,
	6: 5,
	7: 5,
	8: 5,
	9: 5,
}

// map[digit][blink][]numsAtBlink
// TODO: make array instead of map
func buildCache() map[int]map[int]int {
	cache := map[int]map[int]int{}
	cache[0] = map[int]int{}
	stones := []int{0}
	for i := range 5 {
		if i == 0 {
			//cache[0][i] = make([]int, 0)
			cache[0][i] = len(stones)
		} else {
			stones = blink(stones)
			cache[0][i] = len(stones)
		}

	}
	for j := 1; j <= 4; j++ {
		cache[j] = map[int]int{}
		stones = []int{j}
		for i := range 4 {
			if i != 0 {
				stones = blink(stones)
			}
			cache[j][i] = len(stones)
		}
	}
	for j := 5; j <= 7; j++ {
		cache[j] = map[int]int{}
		stones = []int{j}
		for i := range 6 {
			if i != 0 {
				stones = blink(stones)
			}
			cache[j][i] = len(stones)
		}
	}

	cache[8] = map[int]int{
		0: 1,
		1: 1,
		2: 1,
		3: 2,
		4: 4,
		5: 7,
	}

	for j := 9; j < 10; j++ {
		cache[j] = map[int]int{}
		stones = []int{j}
		for i := range 6 {
			if i != 0 {
				stones = blink(stones)
			}
			cache[j][i] = len(stones)
		}
	}
	return cache
}

// num must be single digit
func blinkLength(num int, numBlinks int, cache *map[int]map[int]int) int {
	if numBlinks == 0 {
		return 1
	}
	val, ok := (*cache)[num][numBlinks]
	// in cache, base case
	if ok {
		return val
	}

	totLen := 0
	// too big
	if num >= 10 {
		nextStones := blink([]int{num})
		for _, nextStone := range nextStones {
			nextLength := blinkLength(nextStone, numBlinks-1, cache)
			totLen += nextLength
		}
	} else {
		// not in cache, single digit
		for _, i := range stoneCache[num] {
			if i == 16192 {
				//fmt.Println("stop")
				nextBlinks := numBlinks - 4
				nextLength := blinkLength(8, nextBlinks, cache)
				totLen += nextLength
				(*cache)[8][nextBlinks] = nextLength
			} else {
				nextBlinks := numBlinks - blinkCache[num]
				nextLength := blinkLength(i, nextBlinks, cache)
				totLen += nextLength
				(*cache)[i][nextBlinks] = nextLength
			}
		}
	}
	return totLen
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	stones := []int{}
	for _, char := range strings.Split(data, " ") {
		num, _ := strconv.Atoi(char)
		stones = append(stones, num)
	}

	cache := buildCache()
	tot := 0
	for _, stone := range stones {
		tot += blinkLength(stone, 25, &cache)
	}
	fmt.Println(tot)
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	stones := []int{}
	for _, char := range strings.Split(data, " ") {
		num, _ := strconv.Atoi(char)
		stones = append(stones, num)
	}

	cache := buildCache()
	tot := 0
	for _, stone := range stones {
		tot += blinkLength(stone, 75, &cache)
	}
	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
