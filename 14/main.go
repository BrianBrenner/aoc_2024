package main

import (
	"fmt"
	"os"
	"strings"
)

var Green = "\033[32m"
var Reset = "\033[0m"

func pythonModulo(a, b int) int {
	mod := a % b
	if (mod < 0 && b > 0) || (mod > 0 && b < 0) {
		mod += b
	}
	return mod
}

// returns the quadrant, 1,2,3, or 4 or 0 if in middle
func getQuadrant(startX, startY, speedX, speedY, maxX, maxY, seconds int) int {
	finalX := startX + (speedX * seconds)
	finalX = pythonModulo(finalX, maxX)
	finalY := startY + (speedY * seconds)
	finalY = pythonModulo(finalY, maxY)

	if finalX < maxX/2 && finalY < maxY/2 {
		return 1
	}
	if finalX < maxX/2 && finalY > maxY/2 {
		return 2
	}
	if finalX > maxX/2 && finalY < maxY/2 {
		return 3
	}
	if finalX > maxX/2 && finalY > maxY/2 {
		return 4
	}

	return 0
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	//maxX, maxY := 11, 7
	raw, _ := os.ReadFile("input.txt")
	maxX, maxY := 101, 103
	data := string(raw)
	guards := strings.Split(data, "\n")
	quadrants := map[int]int{}
	for _, guard := range guards {
		startX, startY, speedX, speedY := 0, 0, 0, 0
		//p=0,4 v=3,-3
		fmt.Sscanf(guard, "p=%d,%d v=%d,%d", &startX, &startY, &speedX, &speedY)
		quadrant := getQuadrant(startX, startY, speedX, speedY, maxX, maxY, 100)
		if quadrant != 0 {
			quadrants[quadrant] = quadrants[quadrant] + 1
		}
	}

	tot := 1
	for _, q := range quadrants {
		tot *= q
	}
	fmt.Println(tot)
}

type Guard struct {
	Xpos   int
	Ypos   int
	Xspeed int
	Yspeed int
}

func moveGuard(guard *Guard, maxX, maxY int) {
	finalX := guard.Xpos + guard.Xspeed
	finalX = pythonModulo(finalX, maxX)
	finalY := guard.Ypos + guard.Yspeed
	finalY = pythonModulo(finalY, maxY)

	guard.Xpos = finalX
	guard.Ypos = finalY
}

func printGuards(guards []*Guard, maxX, maxY int) {
	grid := [][]bool{}
	for range maxY {
		row := make([]bool, maxX)
		grid = append(grid, row)
	}
	for _, g := range guards {
		grid[g.Ypos][g.Xpos] = true
	}

	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] {
				fmt.Print(Green + "X" + Reset)
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("")
	}
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	//maxX, maxY := 11, 7
	raw, _ := os.ReadFile("input.txt")
	maxX, maxY := 101, 103
	data := string(raw)
	guardRaw := strings.Split(data, "\n")
	// [[xPos,yPos, ],...]
	guards := []*Guard{}
	for _, g := range guardRaw {
		startX, startY, speedX, speedY := 0, 0, 0, 0
		//p=0,4 v=3,-3
		fmt.Sscanf(g, "p=%d,%d v=%d,%d", &startX, &startY, &speedX, &speedY)
		guards = append(guards, &Guard{
			Xpos:   startX,
			Ypos:   startY,
			Xspeed: speedX,
			Yspeed: speedY,
		})
	}

	for i := range 100 {
		fmt.Println(i)
		printGuards(guards, maxX, maxY)
		for _, g := range guards {
			moveGuard(g, maxX, maxY)
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")

		fmt.Println("")
		fmt.Println("")
		fmt.Println("")

	}

}

func main() {
	//part1()
	part2()
}
