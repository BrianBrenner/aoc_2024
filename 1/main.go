package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1() {
	//file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	list1 := []int{}
	list2 := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	slices.Sort(list1)
	slices.Sort(list2)
	tot := 0.0
	for i := range len(list1) {
		tot += math.Abs(float64(list1[i] - list2[i]))
	}
	fmt.Println(int(tot))
}

func part2() {
	//file, err := os.Open("test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	l1 := []string{}
	counter2 := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		l1 = append(l1, nums[0])
		counter2[nums[1]] += 1
	}

	tot := 0
	for _, v := range l1 {
		kInt, _ := strconv.Atoi(v)
		val := kInt * counter2[v]
		tot += val
	}
	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
