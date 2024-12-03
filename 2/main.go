package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// return 1 if safe, 0 otherwise
func isSafe(nums []int) int {
	prev := -1
	for _, num := range nums {
		if prev == -1 {
			prev = num
			continue
		}
		diff := num - prev
		if diff < 1 || diff > 3 {
			return 0
		}
		prev = num
	}

	return 1
}

func part1() {
	//file, _ := os.Open("test.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numSafe := 0
	for scanner.Scan() {
		line := scanner.Text()
		numString := strings.Fields(line)
		nums := []int{}
		for _, s := range numString {
			val, _ := strconv.Atoi(s)
			nums = append(nums, val)
		}

		numSafe += isSafe(nums)
		slices.Reverse(nums)
		numSafe += isSafe(nums)
	}

	fmt.Println(numSafe)
}

// return 1 if safe, 0 otherwise
func isSafe2(nums []int, skipped int) int {
	for i := range len(nums) - 1 {
		cur := nums[i]
		next := nums[i+1]
		diff := next - cur
		if diff < 1 || diff > 3 {
			skipped += 1
			if skipped > 1 {
				return 0
			}
			withoutFirst := slices.Concat(nums[0:i], nums[i+1:])
			withoutSecond := []int{}
			if i+2 < len(nums) {
				withoutSecond = slices.Concat(nums[0:i+1], nums[i+2:])
			} else {
				withoutSecond = nums[0 : len(nums)-1]
			}
			one := isSafe2(withoutFirst, 1)
			two := isSafe2(withoutSecond, 1)
			if one == 0 && two == 0 {
				return 0
			}
			return 1
		}
	}

	return 1
}

func part2() {
	//file, _ := os.Open("test.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numSafe := 0
	for scanner.Scan() {
		line := scanner.Text()
		numString := strings.Fields(line)
		nums := []int{}
		for _, s := range numString {
			val, _ := strconv.Atoi(s)
			nums = append(nums, val)
		}

		numSafe += isSafe2(nums, 0)
		slices.Reverse(nums)
		numSafe += isSafe2(nums, 0)
	}

	fmt.Println(numSafe)
}

func main() {
	part1()
	part2()
}
