package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO: cleanup and simplify

func helper(target int, nums []int, i int, cur int, op string) bool {
	if i == len(nums) {
		return cur == target
	}
	if op == "+" {
		cur += nums[i]
	} else if op == "*" {
		cur *= nums[i]
	} else {
		fmt.Println("bad")
	}
	return helper(target, nums, i+1, cur, "*") || helper(target, nums, i+1, cur, "+")
}

func helper2(target int, nums []int, i int, cur int, op string) bool {
	if i == len(nums) {
		return cur == target
	}
	if op == "+" {
		cur += nums[i]
	} else if op == "*" {
		cur *= nums[i]
	} else if op == "||" {
		cur, _ = strconv.Atoi(strconv.Itoa(cur) + strconv.Itoa(nums[i]))
	} else {
		fmt.Println("bad")
	}
	return helper2(target, nums, i+1, cur, "*") || helper2(target, nums, i+1, cur, "+") || helper2(target, nums, i+1, cur, "||")
}

func getScore(line string) int {
	split := strings.Split(line, ": ")
	target, _ := strconv.Atoi(split[0])
	numsRaw := strings.Split(split[1], " ")
	nums := []int{}
	for _, n := range numsRaw {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}

	if helper(target, nums, 1, nums[0], "+") || helper(target, nums, 1, nums[0], "*") {
		return target
	}
	return 0
}

func getScore2(line string) int {
	split := strings.Split(line, ": ")
	target, _ := strconv.Atoi(split[0])
	numsRaw := strings.Split(split[1], " ")
	nums := []int{}
	for _, n := range numsRaw {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}

	if helper2(target, nums, 1, nums[0], "+") || helper2(target, nums, 1, nums[0], "*") || helper2(target, nums, 1, nums[0], "||") {
		return target
	}
	return 0
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	lines := strings.Split(data, "\n")

	tot := 0
	for _, line := range lines {
		tot += getScore(line)
	}
	fmt.Println(tot)
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	lines := strings.Split(data, "\n")

	tot := 0
	for _, line := range lines {
		tot += getScore2(line)
	}
	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
