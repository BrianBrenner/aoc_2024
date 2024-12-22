package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func findCheapest(presses [][]int) int {
	if len(presses) == 0 {
		return 0
	}

	minCost := presses[0][0]*3 + presses[0][1]*1
	for _, press := range presses[1:] {
		cost := press[0]*3 + press[1]*3
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func getTokens(game string) int {
	lines := strings.Split(game, "\n")
	xA := 0
	yA := 0
	//Button A: X+94, Y+34
	fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &xA, &yA)
	xB := 0
	yB := 0
	fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &xB, &yB)
	// Prize: X=8400, Y=5400
	prizeX := 0
	prizeY := 0
	fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)

	// contains valid combinations of buttons
	// [[A_presses1, B_presses1], [A_presses2, B_presses2]....]
	valid := [][]int{}
	for a := range 100 {
		for b := range 100 {
			if a*xA+b*xB == prizeX && a*yA+b*yB == prizeY {
				valid = append(valid, []int{a, b})
			}
		}
	}

	return findCheapest(valid)
}

func getTokens2(game string) int {
	lines := strings.Split(game, "\n")
	xA := 0
	yA := 0
	//Button A: X+94, Y+34
	fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &xA, &yA)
	xB := 0
	yB := 0
	fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &xB, &yB)
	// Prize: X=8400, Y=5400
	prizeX := 0
	prizeY := 0
	fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)
	prizeX += 10000000000000
	prizeY += 10000000000000

	// solve the system of equations. luckily, seems like there are no cases with infinitely many solutions
	b := float64(prizeY*xA-yA*prizeX) / float64(xA*yB-yA*xB)
	if b != math.Floor(b) {
		return 0
	}
	a := (float64(prizeX) - b*float64(xB)) / float64(xA)
	if a != math.Floor(a) {
		return 0
	}

	return int(3*a + b)
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	games := strings.Split(data, "\n\n")
	tot := 0
	for _, game := range games {
		tot += getTokens(game)
	}
	fmt.Println(tot)
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	games := strings.Split(data, "\n\n")
	tot := 0
	for _, game := range games {
		tot += getTokens2(game)
	}
	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
