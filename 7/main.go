package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func helper(target int, nums []int, i int, cur int, op string, validOps []string) bool {
	if cur > target {
		return false
	}
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
		fmt.Println("bad input")
	}

	out := false
	for _, nextOp := range validOps {
		out = out || helper(target, nums, i+1, cur, nextOp, validOps)
	}
	return out
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	lines := strings.Split(data, "\n")

	tot := 0
	for _, line := range lines {
		split := strings.Split(line, ": ")
		target, _ := strconv.Atoi(split[0])
		numsRaw := strings.Split(split[1], " ")
		nums := []int{}
		for _, n := range numsRaw {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		validOps := []string{"+", "*"}
		if helper(target, nums, 1, nums[0], "+", validOps) || helper(target, nums, 1, nums[0], "*", validOps) {
			tot += target
		}

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
		split := strings.Split(line, ": ")
		target, _ := strconv.Atoi(split[0])
		numsRaw := strings.Split(split[1], " ")
		nums := []int{}
		for _, n := range numsRaw {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		validOps := []string{"+", "*", "||"}
		if helper(target, nums, 1, nums[0], "+", validOps) || helper(target, nums, 1, nums[0], "*", validOps) || helper(target, nums, 1, nums[0], "||", validOps) {
			tot += target
		}
	}
	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
