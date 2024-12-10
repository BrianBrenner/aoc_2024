package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func compact(file []string) []string {
	start := 0
	end := len(file) - 1
	for start < end {
		if file[start] == "." {
			for start < end {
				if file[end] == "." {
					end--
				} else {
					break
				}
			}
			file[start] = file[end]
			file[end] = "."
			end--
		}
		start++
	}

	return file
}

func compact2(file []filePiece) []string {
	//out := make([]string, len(file))

	for end := len(file) - 1; end >= 0; end-- {
		if file[end].Val != "." {
			for i := 0; i < end; i++ {
				// left most space has room
				if file[i].Val == "." && file[i].Len >= file[end].Len {
					leftover := file[i].Len - file[end].Len
					file[i].Val = file[end].Val
					file[i].Len = file[end].Len
					file[end].Val = "."
					if leftover > 0 {
						slices.Insert(file, i+1, filePiece{
							Val: ".",
							Len: leftover,
						})
						// doesn't change answer, but I think this would be needed?
						end++
					}
					break
				}
			}
		}
	}

	out := []string{}
	for _, piece := range file {
		for range piece.Len {
			out = append(out, piece.Val)
		}
	}

	return out
}

func checksum(file []string) int {
	tot := 0
	for pos, char := range file {
		val, _ := strconv.Atoi(string(char))
		tot += val * pos
	}
	return tot
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	file := []string{}
	id := 0
	for i, char := range data {
		num, _ := strconv.Atoi(string(char))
		// file
		if i%2 == 0 {
			for range num {
				file = append(file, strconv.Itoa(id))
			}
			id++
		} else { // freespace
			for range num {
				file = append(file, ".")
			}
		}
	}
	fmt.Println(checksum(compact(file)))
}

type filePiece struct {
	Val string
	Len int
}

func part2() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	file := []filePiece{}
	id := 0
	for i, char := range data {
		num, _ := strconv.Atoi(string(char))
		// file
		if i%2 == 0 {
			piece := filePiece{
				Val: strconv.Itoa(id),
				Len: num,
			}
			file = append(file, piece)
			id++
		} else { // freespace
			piece := filePiece{
				Val: ".",
				Len: num,
			}
			file = append(file, piece)
		}
	}
	fmt.Println(checksum(compact2(file)))
}

func main() {
	part1()
	part2()
}
