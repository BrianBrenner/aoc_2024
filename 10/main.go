package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// finishes uses the coords as a key, ie "r,c" ex: "1,3". if
func dfs(grid [][]int, r int, c int, prev int, finishes *map[string]bool) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}
	if grid[r][c] != prev+1 {
		return
	}
	if grid[r][c] == 9 {
		coordKey := strconv.Itoa(r) + "," + strconv.Itoa(c)
		(*finishes)[coordKey] = true
	}
	dfs(grid, r-1, c, grid[r][c], finishes)
	dfs(grid, r+1, c, grid[r][c], finishes)
	dfs(grid, r, c+1, grid[r][c], finishes)
	dfs(grid, r, c-1, grid[r][c], finishes)
}

func dfs2(grid [][]int, r int, c int, prev int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return 0
	}
	if grid[r][c] != prev+1 {
		return 0
	}
	if grid[r][c] == 9 {
		return 1
	}
	return dfs2(grid, r-1, c, grid[r][c]) + dfs2(grid, r+1, c, grid[r][c]) + dfs2(grid, r, c+1, grid[r][c]) + dfs2(grid, r, c-1, grid[r][c])
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	grid := [][]int{}
	for _, line := range strings.Split(data, "\n") {
		row := []int{}
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	tot := 0
	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == 0 {
				finishes := map[string]bool{}
				dfs(grid, r, c, -1, &finishes)
				tot += len(finishes)
			}
		}
	}
	fmt.Println(tot)
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	grid := [][]int{}
	for _, line := range strings.Split(data, "\n") {
		row := []int{}
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	tot := 0
	for r := range len(grid) {
		for c := range len(grid[0]) {
			if grid[r][c] == 0 {
				tot += dfs2(grid, r, c, -1)
			}
		}
	}
	fmt.Println(tot)
}

func main() {
	part1()
	part2()
}
