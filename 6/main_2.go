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

func emptyVisited(n int) [][]map[string]bool {
	out := make([][]map[string]bool, n) // create the outer slice with space for 'rows' inner slices
	for i := range n {
		out[i] = make([]map[string]bool, n) // create each inner slice with 'cols' elements
		for j := range n {
			out[i][j] = make(map[string]bool)
		}
	}
	return out
}

func copyVisited(visited [][][]string) [][][]string {
	// Create a new slice with the same length as the original
	copied := make([][][]string, len(visited))

	for i, subArray := range visited {
		// Create a new slice for each 2D sub-array
		copied[i] = make([][]string, len(subArray))

		for j, innerArray := range subArray {
			// Create a new slice for each 1D inner-array
			copied[i][j] = make([]string, len(innerArray))

			// Copy the strings from the original inner-array
			copy(copied[i][j], innerArray)
		}
	}

	return copied
}

// updates grid in place by moving the arrow one step, updates visited in place with position at start of step
// returns r, c, bool where bool indicates out of bounds if were to take a step, ie you are on the very edge
func step(r, c int, gridPtr *[][]string, visited *[][]map[string]bool) (int, int, bool) {
	grid := *gridPtr
	curDir := grid[r][c]
	// update visited
	(*visited)[r][c][curDir] = true
	dirCoord := dirCoords[curDir]
	nextR, nextC := r+dirCoord[0], c+dirCoord[1]
	nextInBound := nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0])
	if !nextInBound {
		return r, c, true
	}
	// check for turn
	if grid[nextR][nextC] == "#" {
		// stay in place and turn
		turnedDir := nextDirs[curDir]
		grid[r][c] = turnedDir
		// since turned have to check if next step out of bounds
		dirCoord = dirCoords[turnedDir]
		nextR, nextC = r+dirCoord[0], c+dirCoord[1]
		nextInBound = nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0])
		if !nextInBound {
			return r, c, true
		}
		return r, c, false
	}
	// step in current direction
	grid[r][c] = "."
	grid[nextR][nextC] = curDir

	return nextR, nextC, false
}

// takes in a grid with a start location. traverses the grid till it exits or get in a loop.
// returns true if has a loop
func traverse(r, c int, origGrid [][]string) (hasLoop bool) {
	grid := copyGrid(origGrid)
	visited := emptyVisited(len(origGrid))
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
	//r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func part2() {
	raw, _ := os.ReadFile("input.txt")
	//raw, _ := os.ReadFile("test.txt")
	data := string(raw)

	grid := [][]string{}
	for _, row := range strings.Split(data, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}

	startR := -1
	startC := -1
	for i := range len(grid) {
		for j := range len(grid[0]) {
			// find start
			if grid[i][j] == "^" {
				startR = i
				startC = j
				break
			}
		}
	}

	count := 0
	// visited doesn't matter, we don't check it here
	visited := emptyVisited(len(grid))
	gridCopy := copyGrid(grid)
	r := startR
	c := startC
	for {
		done := false
		r, c, done = step(r, c, &gridCopy, &visited)
		if done {
			fmt.Println(count)
			return
		}

		// imagine there is an obstacle in front of us
		withObstacle := copyGrid(gridCopy)
		curDir := withObstacle[r][c]
		// place obstacle, we know one can exist there since done is false
		dirCoord := dirCoords[curDir]
		obstacleR, obstacleC := r+dirCoord[0], c+dirCoord[1]
		nextInBound := obstacleR >= 0 && obstacleR < len(grid) && obstacleC >= 0 && obstacleC < len(grid[0])
		if !nextInBound {
			fmt.Println(count)
			return
		}
		withObstacle[obstacleR][obstacleC] = "#"

		hasLoop := traverse(r, c, withObstacle)
		if hasLoop {
			count += 1
		}
	}
}

func main() {
	//part1()
	part2()
}
