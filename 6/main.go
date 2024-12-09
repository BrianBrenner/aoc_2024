package main

import (
	"fmt"
	"os"
	"strings"
)

var dirCoords = map[string][]int{
	"^": []int{-1, 0},
	"v": []int{1, 0},
	">": []int{0, 1},
	"<": []int{0, -1},
}

var nextDirs = map[string]string{
	"^": ">",
	"v": "<",
	">": "v",
	"<": "^",
}

func copyGrid(grid [][]string) [][]string {
	c := [][]string{}
	for i := range len(grid) {
		row := []string{}
		for j := range len(grid[0]) {
			row = append(row, grid[i][j])
		}
		c = append(c, row)
	}
	return c
}

// visited contains which direction we moved in at each coordinate
func emptyVisited(n int) [][]map[string]bool {
	out := make([][]map[string]bool, n)
	for i := range n {
		out[i] = make([]map[string]bool, n)
		for j := range n {
			out[i][j] = make(map[string]bool)
		}
	}
	return out
}

// updates grid in place by moving the arrow one step, updates visited in place with position at start of step
// returns r, c, bool where bool indicates out of bounds if were to take a step, ie you are on the very edge
func step(r, c int, gridPtr *[][]string, visited *[][]map[string]bool) (int, int, bool) {
	grid := *gridPtr
	curDir := grid[r][c]
	// make location as visited with the current direction then take a step if valid
	(*visited)[r][c][curDir] = true
	nextR, nextC := r+dirCoords[curDir][0], c+dirCoords[curDir][1]
	nextInBounds := nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0])
	if !nextInBounds {
		return r, c, true
	}
	// check for turn
	if grid[nextR][nextC] == "#" {
		// stay in place and turn
		turnedDir := nextDirs[curDir]
		grid[r][c] = turnedDir
		return r, c, false
	}
	// step in current direction
	grid[r][c] = "."
	grid[nextR][nextC] = curDir

	return nextR, nextC, false
}

// takes in a grid with a start location. traverses the grid till it exits or gets in a loop.
// returns true if has a loop
func traverse(r, c int, origGrid [][]string) bool {
	grid := copyGrid(origGrid)
	visited := emptyVisited(len(origGrid))
	// break if make it out of grid or get in a loop
	for {
		done := false
		r, c, done = step(r, c, &grid, &visited)
		if done {
			return false
		}

		curDir := grid[r][c]
		looped := visited[r][c][curDir]
		if looped {
			return true
		}
	}
}

func part1() {
	raw, _ := os.ReadFile("input.txt")
	//raw, _ := os.ReadFile("test.txt")
	data := string(raw)

	origGrid := [][]string{}
	for _, row := range strings.Split(data, "\n") {
		origGrid = append(origGrid, strings.Split(row, ""))
	}

	// find start
	startR := -1
	startC := -1
	for i := range len(origGrid) {
		for j := range len(origGrid[0]) {
			if origGrid[i][j] == "^" {
				startR = i
				startC = j
				break
			}
		}
	}

	visited := emptyVisited(len(origGrid))
	grid := copyGrid(origGrid)
	r := startR
	c := startC
	for {
		// walk through the guards path and at each step imagine there is an obstacle in front of us and see if it would cause
		// a loop
		done := false
		r, c, done = step(r, c, &grid, &visited)
		if done {
			count := 0
			for i := range len(visited) {
				for j := range len(visited[0]) {
					if len(visited[i][j]) > 0 {
						count++
					}
				}
			}
			fmt.Println(count)

			return
		}
	}
}

func part2() {
	raw, _ := os.ReadFile("input.txt")
	//raw, _ := os.ReadFile("test.txt")
	data := string(raw)

	origGrid := [][]string{}
	for _, row := range strings.Split(data, "\n") {
		origGrid = append(origGrid, strings.Split(row, ""))
	}

	// find start
	startR := -1
	startC := -1
	for i := range len(origGrid) {
		for j := range len(origGrid[0]) {
			if origGrid[i][j] == "^" {
				startR = i
				startC = j
				break
			}
		}
	}

	visited := emptyVisited(len(origGrid))
	grid := copyGrid(origGrid)
	count := 0
	r := startR
	c := startC
	for {
		// walk through the guards path and at each step imagine there is an obstacle in front of us and see if it would cause
		// a loop
		done := false
		r, c, done = step(r, c, &grid, &visited)
		if done {
			fmt.Println(count)
			return
		}

		curDir := grid[r][c]
		withObstacle := copyGrid(origGrid)
		// move start to current position
		withObstacle[startR][startC] = "."
		withObstacle[r][c] = curDir

		// make sure obstacle goes in a valid spot
		dirCoord := dirCoords[curDir]
		obstacleR, obstacleC := r+dirCoord[0], c+dirCoord[1]
		nextInBounds := obstacleR >= 0 && obstacleR < len(origGrid) && obstacleC >= 0 && obstacleC < len(origGrid[0])
		if !nextInBounds {
			fmt.Println(count)
			return
		}
		// this means we already visited where the obstacle would go so we would have ran into ito earlier so skip it
		if len(visited[obstacleR][obstacleC]) > 0 {
			continue
		}
		// place obstacle, and check if we started from there would we get in a loop
		withObstacle[obstacleR][obstacleC] = "#"
		hasLoop := traverse(r, c, withObstacle)
		if hasLoop {
			count += 1
		}
	}
}

func main() {
	part1()
	part2()
}
