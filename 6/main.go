package main

import (
	"fmt"
	"os"
	"slices"
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

func move(r, c int, grid, visited [][]string) (int, int, [][]string, [][]string) {
	visited[r][c] = "X"
	curDir := grid[r][c]
	dirCoord := dirCoords[curDir]
	nextR := r + dirCoord[0]
	nextC := c + dirCoord[1]

	if !(nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0])) {
		return nextR, nextC, grid, visited
	}

	if grid[nextR][nextC] == "#" { // need to turn
		nextDir := nextDirs[curDir]
		grid[r][c] = nextDir
	} else { // stay same direction
		grid[r][c] = "."
		r = nextR
		c = nextC
		grid[r][c] = curDir
	}

	return r, c, grid, visited
}

func move2(r, c int, grid [][]string, visitedDirs [][][]string) (int, int, [][]string, [][][]string, bool) {
	curDir := grid[r][c]
	dirCoord := dirCoords[curDir]
	visitedDirs[r][c] = append(visitedDirs[r][c], curDir)
	nextR := r + dirCoord[0]
	nextC := c + dirCoord[1]

	if !(nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0])) {
		return nextR, nextC, grid, visitedDirs, false
	}

	couldLoop := false
	nextDir := ""
	if grid[nextR][nextC] == "#" { // need to turn
		nextDir = nextDirs[curDir]
		nextDirCoord := dirCoords[nextDir]
		nextR = r + nextDirCoord[0]
		nextC = c + nextDirCoord[1]

		grid[r][c] = nextDir
	} else { // stay same direction
		nextDir = curDir
		// check if a loop is possible by imagining we need to turn and see if we follow that new direction we see a visitedDir
		// that matches the new direction
		loopDir := nextDirs[curDir]
		loopDirCoord := dirCoords[loopDir]
		loopR := r + loopDirCoord[0]
		loopC := c + loopDirCoord[1]

		couldLoop = checkLoop(loopR, loopC, loopDir, grid, visitedDirs)
	}

	grid[r][c] = "."
	r = nextR
	c = nextC
	grid[r][c] = nextDir

	return r, c, grid, visitedDirs, couldLoop
}

func checkLoop(r, c int, loopDir string, grid [][]string, visitedDirs [][][]string) bool {
	if loopDir == "^" {
		for i := r; i >= 0; i-- {
			if slices.Contains(visitedDirs[i][c], loopDir) {
				return true
			}
		}
		return false
	}
	if loopDir == "v" {
		for i := r; i < len(visitedDirs); i++ {
			if slices.Contains(visitedDirs[i][c], loopDir) {
				return true
			}
		}
		return false
	}
	if loopDir == ">" {
		for i := c; i < len(visitedDirs[0]); i++ {
			if slices.Contains(visitedDirs[r][i], loopDir) {
				return true
			}
		}
		return false
	}
	if loopDir == "<" {
		for i := c; i >= 0; i-- {
			if slices.Contains(visitedDirs[r][i], loopDir) {
				return true
			}
		}
		return false
	}
	return false
}

//func checkLoop(r, c int, grid [][]string, visitedDirs [][][]string) bool {
//	for r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
//		r, c, grid, visitedDirs, hasLoop = move2(r, c, grid, visitedDirs)
//		if hasLoop {
//			return true
//		}
//	}
//	return false
//}

func part1() {
	raw, _ := os.ReadFile("input.txt")
	//raw, _ := os.ReadFile("test.txt")
	data := string(raw)

	grid := [][]string{}
	visited := [][]string{}
	for _, row := range strings.Split(data, "\n") {
		grid = append(grid, strings.Split(row, ""))
		visited = append(visited, strings.Split(row, ""))
	}

	r := -1
	c := -1
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == "^" {
				r = i
				c = j
				break
			}
		}
	}

	for r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
		r, c, grid, visited = move(r, c, grid, visited)
	}

	tot := 0
	for i := range len(visited) {
		for j := range len(visited[0]) {
			if visited[i][j] == "X" {
				tot += 1
			}
		}
	}
	fmt.Println(tot)
}

func part2() {
	//raw, _ := os.ReadFile("input.txt")
	raw, _ := os.ReadFile("test.txt")
	data := string(raw)

	grid := [][]string{}
	visitedDirs := [][][]string{}
	for _, row := range strings.Split(data, "\n") {
		grid = append(grid, strings.Split(row, ""))
		dirRow := [][]string{}
		for range len(strings.Split(row, "")) {
			dirRow = append(dirRow, []string{})
		}
		visitedDirs = append(visitedDirs, dirRow)
	}

	r := -1
	c := -1
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == "^" {
				r = i
				c = j
				break
			}
		}
	}

	count := 0
	couldLoop := false
	for r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
		r, c, grid, visitedDirs, couldLoop = move2(r, c, grid, visitedDirs)
		if couldLoop {
			count += 1
		}
	}
	fmt.Println(count)
}

func main() {
	//part1()
	part2()
}
