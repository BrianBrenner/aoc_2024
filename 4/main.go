package main

import (
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getCount(strs []string) int {
	count := 0
	for _, str := range strs {
		count += strings.Count(str, "XMAS")
		count += strings.Count(reverse(str), "XMAS")
	}

	return count
}

func getDiag(r, c int, grid []string) string {
	diag := ""
	for r < len(grid) && c < len(grid[0]) {
		diag += string(grid[r][c])
		r += 1
		c += 1
	}
	return diag
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	lines := strings.Split(data, "\n")

	count := 0
	// horizontal
	count += getCount(lines)

	// vertical
	verts := []string{}
	for c := range len(lines[0]) {
		str := ""
		for r := range len(lines) {
			str += string(lines[r][c])
		}
		verts = append(verts, str)
	}
	count += getCount(verts)

	diags := []string{}
	for c := range len(lines[0]) {
		diags = append(diags, getDiag(0, c, lines))
	}
	for r := 1; r < len(lines); r++ {
		diags = append(diags, getDiag(r, 0, lines))
	}

	reversedLines := []string{}
	for _, line := range lines {
		reversedLines = append(reversedLines, reverse(line))
	}
	for c := range len(reversedLines[0]) {
		diags = append(diags, getDiag(0, c, reversedLines))
	}
	for r := 1; r < len(reversedLines); r++ {
		diags = append(diags, getDiag(r, 0, reversedLines))
	}

	count += getCount(diags)

	fmt.Println(count)
}

func getBoxes(lines []string) [][]string {
	boxes := [][]string{}
	for r := 0; r < len(lines)-2; r++ {
		for c := 0; c < len(lines[0])-2; c++ {
			box := []string{lines[r][c : c+3], lines[r+1][c : c+3], lines[r+2][c : c+3]}
			boxes = append(boxes, box)
		}
	}

	return boxes
}

func isXmas(box []string) bool {
	diag1 := string(box[0][0]) + string(box[1][1]) + string(box[2][2])
	diag2 := string(box[0][2]) + string(box[1][1]) + string(box[2][0])
	return (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM")
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	lines := strings.Split(data, "\n")

	boxes := getBoxes(lines)
	count := 0
	for _, box := range boxes {
		if isXmas(box) {
			count += 1
		}
	}
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
