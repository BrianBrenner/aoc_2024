package main

import (
	"fmt"
	"os"
	"slices"
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

func dfs(grid [][]string, cur Step, finishes *[]int) {
	// reached end
	if grid[cur.R][cur.C] == "E" {
		*finishes = append(*finishes, cur.Points)
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
				dfs(grid, next, finishes)
			}
		}
	}
}

// TODO: maybe do a weighted bfs? like if you have to turn, put it at the back of the queue, but if you are going
// straight put at the front?

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

	//minPoints := -1
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		// reached end
		if grid[cur.R][cur.C] == "E" {
			return cur.Points
			//if minPoints == -1 {
			//	minPoints = cur.Points
			//} else if cur.Points < minPoints {
			//	minPoints = cur.Points
			//}
			//continue
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
					next := Step{
						R:       nextR,
						C:       nextC,
						Dir:     dir,
						Visited: &nextVisited,
					}

					// i think we don't have to worry about turning 180 degrees, because that would have been visited
					if cur.Dir != dir {
						next.Points = cur.Points + 1001
						q = append(q, next)
					} else {
						next.Points = cur.Points + 1
						q = append([]Step{next}, q...)
					}
				}
			}
		}
	}

	return -1
}

func helper(grid [][]string, startR, startC int) int {
	v := createVisited(len(grid), len(grid[0]))
	v[startR][startC] = true
	start := Step{
		R:       startR,
		C:       startC,
		Dir:     complex(1, 0),
		Points:  0,
		Visited: &v,
	}
	finishes := []int{}
	dfs(grid, start, &finishes)
	return slices.Min(finishes)
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
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
	//fmt.Println(helper(grid, r, c))
}

func main() {
	part1()
}
