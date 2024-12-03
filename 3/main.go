package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	fmt.Println(getCount(data))
}

func getCount(str string) int {
	rx := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
	matches := rx.FindAllString(str, -1)
	tot := 0
	for _, match := range matches {
		var val1 int
		var val2 int
		_, err := fmt.Sscanf(match, "mul(%d,%d)", &val1, &val2)
		if err != nil {
			log.Fatal(err)
		}
		tot += val1 * val2
	}

	return tot
}

func part2() {
	//raw, _ := os.ReadFile("test2.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)

	donts := strings.Split(data, "don't")
	// we start enabled
	tot := getCount(donts[0])
	for _, dont := range donts[1:] {
		dos := strings.Split(dont, "do()")
		if len(dos) < 2 {
			continue
		}
		for _, do := range dos[1:] {
			tot += getCount(do)
		}
	}

	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
