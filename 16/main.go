package main

import (
	"fmt"
	"os"
	"strings"
)

type Step struct {
	R       int
	C       int
	Dir     complex64
	Points  int
	Visited *[][]bool
}

// use complex numbers for directions
var dirs = []complex64{complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)}

func createVisited(rSize, cSize int) [][]bool {
	grid := [][]bool{}
	for range rSize {
		row := make([]bool, cSize)
		grid = append(grid, row)
	}
	return grid
}

func printVisited(grid [][]bool) {
	for r := range len(grid) {
		for c := range len(grid[0]) {
			val := grid[r][c]
			if val {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func cloneVisited(v [][]bool) [][]bool {
	cloned := make([][]bool, len(v))
	for i := range v {
		cloned[i] = make([]bool, len(v[i]))
		copy(cloned[i], v[i])
	}
	return cloned
}

// i think maybe use dfs, or seperate visited array for each iteration. maybe look at day 6
func bfs(grid [][]string, startR, startC int) int {
	q := []Step{}
	v := createVisited(len(grid), len(grid[0]))
	v[startR][startC] = true
	start := Step{
		R:       startR,
		C:       startC,
		Dir:     complex(1, 0),
		Points:  0,
		Visited: &v,
	}
	q = append(q, start)

	minPoints := -1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		// reached end
		if grid[cur.R][cur.C] == "E" {
			if minPoints == -1 {
				minPoints = cur.Points
			} else if cur.Points < minPoints {
				minPoints = cur.Points
			}
			continue
		}
		for _, dir := range dirs {
			nextR := cur.R + int(imag(dir))
			nextC := cur.C + int(real(dir))
			// in bounds
			if nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid) {
				// valid and not visited
				if grid[nextR][nextC] != "#" && !(*cur.Visited)[nextR][nextC] {
					nextVisited := cloneVisited(*cur.Visited)
					nextVisited[nextR][nextC] = true
					nextPoints := cur.Points + 1
					// i think we don't have to worry about turning 180 degrees, because that would have been visited
					if cur.Dir != dir {
						nextPoints = cur.Points + 1001
					}
					next := Step{
						R:       nextR,
						C:       nextC,
						Dir:     dir,
						Points:  nextPoints,
						Visited: &nextVisited,
					}
					q = append(q, next)
				}
			}
		}
	}

	return minPoints
}

func part1() {
	raw, _ := os.ReadFile("test.txt")
	//raw, _ := os.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")
	grid := [][]string{}
	for _, line := range data {
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	r := 0
	c := 0
	// find start
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == "S" {
				r = i
				c = j
			}
		}
	}

	fmt.Println(bfs(grid, r, c))
}

func main() {
	part1()
}
