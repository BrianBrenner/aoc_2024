package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createGrid(bytes [][]int, gridSize int) [][]string {
	grid := [][]string{}
	for range gridSize {
		row := make([]string, gridSize)
		grid = append(grid, row)
	}

	for _, b := range bytes {
		r, c := b[0], b[1]
		grid[r][c] = "#"
	}

	return grid
}

func createVisited(gridSize int) [][]bool {
	grid := [][]bool{}
	for range gridSize {
		row := make([]bool, gridSize)
		grid = append(grid, row)
	}
	return grid
}

func cloneVisited(v [][]bool) [][]bool {
	cloned := make([][]bool, len(v))
	for i := range v {
		cloned[i] = make([]bool, len(v[i]))
		copy(cloned[i], v[i])
	}
	return cloned
}

func printGrid(grid [][]string) {
	for r := range len(grid) {
		for c := range len(grid[0]) {
			val := grid[r][c]
			if val == "" {
				fmt.Print(".")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Print("\n")
	}
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

type Step struct {
	R       int
	C       int
	Steps   int
	Visited *[][]bool
}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func bfs(grid [][]string) Step {
	q := []Step{}
	v := createVisited(len(grid))
	v[0][0] = true
	start := Step{
		R:       0,
		C:       0,
		Steps:   0,
		Visited: &v,
	}
	q = append(q, start)

	for len(q) > 0 {
		cur := q[0]
		if cur.R == len(grid)-1 && cur.C == len(grid)-1 {
			return cur
		}
		q = q[1:]
		for _, dir := range dirs {
			r, c := dir[0], dir[1]
			nextR := cur.R + r
			nextC := cur.C + c
			// in bounds
			if nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid) {
				// valid and not visited
				if grid[nextR][nextC] != "#" && !(*cur.Visited)[nextR][nextC] {
					(*cur.Visited)[nextR][nextC] = true
					next := Step{
						R:       nextR,
						C:       nextC,
						Steps:   cur.Steps + 1,
						Visited: cur.Visited,
					}
					q = append(q, next)
				}
			}
		}
	}

	return Step{}
}

func countSteps(grid [][]string) int {
	end := bfs(grid)
	return end.Steps
}

func findFirstBlock(grid [][]string, bytes [][]int) []int {
	for _, b := range bytes {
		r, c := b[0], b[1]
		grid[r][c] = "#"

		end := bfs(grid)
		if end.Steps == 0 {
			return []int{c, r}
		}
	}
	return []int{}
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	//gridSize := 7
	raw, _ := os.ReadFile("input.txt")
	gridSize := 71
	data := string(raw)
	rawBytes := strings.Split(data, "\n")
	// stores indexes [[r,c],...]
	bytes := [][]int{}
	for _, b := range rawBytes {
		coords := strings.Split(b, ",")
		r, _ := strconv.Atoi(coords[1])
		c, _ := strconv.Atoi(coords[0])
		bytes = append(bytes, []int{r, c})
	}
	grid := createGrid(bytes[:1024], gridSize)
	//grid := createGrid(bytes[:12], gridSize)
	fmt.Println(countSteps(grid))
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	//gridSize := 7
	raw, _ := os.ReadFile("input.txt")
	gridSize := 71
	data := string(raw)
	rawBytes := strings.Split(data, "\n")
	// stores indexes [[r,c],...]
	bytes := [][]int{}
	for _, b := range rawBytes {
		coords := strings.Split(b, ",")
		r, _ := strconv.Atoi(coords[1])
		c, _ := strconv.Atoi(coords[0])
		bytes = append(bytes, []int{r, c})
	}
	//grid := createGrid(bytes[:12], gridSize)
	grid := createGrid(bytes[:1024], gridSize)
	fmt.Println(findFirstBlock(grid, bytes[1024:]))
}

func main() {
	part1()
	part2()
}
