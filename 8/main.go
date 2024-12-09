package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	R int
	C int
}

func findNodes(coords []coord, nodes *[][]bool) {
	for i := range len(coords) {
		for j := i + 1; j < len(coords); j++ {
			coord1 := coords[i]
			coord2 := coords[j]
			rDiff := coord2.R - coord1.R
			cDiff := coord2.C - coord1.C
			n1 := coord{
				R: coord1.R - rDiff,
				C: coord1.C - cDiff,
			}
			n2 := coord{
				R: coord2.R + rDiff,
				C: coord2.C + cDiff,
			}
			if n1.R >= 0 && n1.R < len(*nodes) && n1.C >= 0 && n1.C < len((*nodes)[0]) {
				(*nodes)[n1.R][n1.C] = true
			}
			if n2.R >= 0 && n2.R < len(*nodes) && n2.C >= 0 && n2.C < len((*nodes)[0]) {
				(*nodes)[n2.R][n2.C] = true
			}
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to simplify the ratio
func simplifyRatio(a, b int) (int, int) {
	divisor := gcd(a, b)
	return a / divisor, b / divisor
}

func findNodes2(coords []coord, nodes *[][]bool) {
	for i := range len(coords) {
		for j := i + 1; j < len(coords); j++ {
			coord1 := coords[i]
			coord2 := coords[j]
			rDiff := coord2.R - coord1.R
			cDiff := coord2.C - coord1.C
			rDiff, cDiff = simplifyRatio(rDiff, cDiff)

			start := coord1
			for start.R >= 0 && start.R < len(*nodes) && start.C >= 0 && start.C < len((*nodes)[0]) {
				(*nodes)[start.R][start.C] = true
				start.R -= rDiff
				start.C -= cDiff
			}
			start = coord1
			for start.R >= 0 && start.R < len(*nodes) && start.C >= 0 && start.C < len((*nodes)[0]) {
				(*nodes)[start.R][start.C] = true
				start.R += rDiff
				start.C += cDiff
			}

			//n1 := coord{
			//	R: coord1.R - rDiff,
			//	C: coord1.C - cDiff,
			//}
			//n2 := coord{
			//	R: coord2.R + rDiff,
			//	C: coord2.C + cDiff,
			//}
			//if n1.R >= 0 && n1.R < len(*nodes) && n1.C >= 0 && n1.C < len((*nodes)[0]) {
			//	(*nodes)[n1.R][n1.C] = true
			//}
			//if n2.R >= 0 && n2.R < len(*nodes) && n2.C >= 0 && n2.C < len((*nodes)[0]) {
			//	(*nodes)[n2.R][n2.C] = true
			//}
		}
	}
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	lines := strings.Split(data, "\n")

	grid := [][]string{}
	nodes := [][]bool{}
	for _, line := range lines {
		row := make([]string, len(line))
		boolRow := make([]bool, len(line))
		for i, char := range strings.Split(line, "") {
			row[i] = char
			boolRow[i] = false
		}
		grid = append(grid, row)
		nodes = append(nodes, boolRow)
	}

	// map from a frequency to coordinates of the towers with that frequency
	freqs := map[string][]coord{}
	for r := range len(grid) {
		for c := range len(grid[0]) {
			val := grid[r][c]
			if val != "." {
				freqs[val] = append(freqs[val], coord{R: r, C: c})
			}

		}
	}

	for _, coords := range freqs {
		findNodes(coords, &nodes)
	}

	count := 0
	for i := range len(nodes) {
		for j := range len(nodes[0]) {
			if nodes[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	lines := strings.Split(data, "\n")

	grid := [][]string{}
	nodes := [][]bool{}
	for _, line := range lines {
		row := make([]string, len(line))
		boolRow := make([]bool, len(line))
		for i, char := range strings.Split(line, "") {
			row[i] = char
			boolRow[i] = false
		}
		grid = append(grid, row)
		nodes = append(nodes, boolRow)
	}

	// map from a frequency to coordinates of the towers with that frequency
	freqs := map[string][]coord{}
	for r := range len(grid) {
		for c := range len(grid[0]) {
			val := grid[r][c]
			if val != "." {
				freqs[val] = append(freqs[val], coord{R: r, C: c})
			}

		}
	}

	for _, coords := range freqs {
		findNodes2(coords, &nodes)
	}

	count := 0
	for i := range len(nodes) {
		for j := range len(nodes[0]) {
			if nodes[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
