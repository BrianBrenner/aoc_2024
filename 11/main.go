package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func blink(nums []int) []int {
	newNums := []int{}
	for _, num := range nums {
		if num == 0 {
			newNums = append(newNums, 1)
			continue
		}
		strNum := strconv.Itoa(num)
		if len(strNum)%2 == 0 {
			l, _ := strconv.Atoi(strNum[:len(strNum)/2])
			r, _ := strconv.Atoi(strNum[len(strNum)/2:])
			newNums = append(newNums, l, r)
			continue
		}
		newNums = append(newNums, num*2024)
	}

	return newNums
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	nums := []int{}
	for _, char := range strings.Split(data, " ") {
		num, _ := strconv.Atoi(char)
		nums = append(nums, num)
	}
	for range 25 {
		nums = blink(nums)
	}
	fmt.Println(len(nums))
}

func main() {
	part1()
}
