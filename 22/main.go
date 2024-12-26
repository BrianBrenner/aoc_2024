package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mix(num int, num2 int) int {
	return num ^ num2
}

func prune(num int) int {
	return num % 16777216
}

func nextNum(num int) int {
	toMix := num * 64
	num = prune(mix(num, toMix))
	toMix = num / 32
	num = prune(mix(num, toMix))
	toMix = num * 2048
	num = prune(mix(num, toMix))

	return num
}

func getScore(num int) int {
	for range 2000 {
		num = nextNum(num)
	}
	return num
}

func getPricesAndChanges(num int) ([]int, []int) {
	prevPrice := num % 10
	prices := []int{prevPrice}
	changes := []int{}
	for range 2000 {
		next := nextNum(num)
		price := next % 10
		changes = append(changes, price-prevPrice)
		prices = append(prices, price)
		num = next
		prevPrice = price
	}
	return prices, changes
}

func getKey(changes []int) string {
	strChange := []string{}
	for _, c := range changes {
		num := strconv.Itoa(c)
		strChange = append(strChange, num)
	}

	return strings.Join(strChange, ",")
}

func updateScores(prices []int, changes []int, scores *map[string]int) {
	seen := map[string]bool{}
	for i := 0; i <= len(changes)-4; i++ {
		changeKey := getKey(changes[i : i+4])
		_, ok := seen[changeKey]
		if ok {
			continue
		}
		seen[changeKey] = true
		(*scores)[changeKey] = (*scores)[changeKey] + prices[i+4]
	}
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")

	tot := 0
	for _, line := range data {
		num, _ := strconv.Atoi(line)
		tot += getScore(num)
	}
	fmt.Println(tot)
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")

	scores := map[string]int{}
	for _, line := range data {
		num, _ := strconv.Atoi(line)
		prices, changes := getPricesAndChanges(num)
		updateScores(prices, changes, &scores)
	}

	maxVal := 0
	for _, v := range scores {
		if v > maxVal {
			maxVal = v
		}
	}
	fmt.Println(maxVal)
}

func main() {
	part1()
	part2()
}
