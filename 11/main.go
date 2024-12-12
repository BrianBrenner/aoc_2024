package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const totalBlinks = 12
// breaks between 11 and 12
const totalBlinks = 11

func blink(stones []Stone) []Stone {
	newStones := []Stone{}
	for _, stone := range stones {
		if stone.Val == 0 {
			newStones = append(newStones, Stone{1, stone.Blinks + 1})
			continue
		}
		strNum := strconv.Itoa(stone.Val)
		if len(strNum)%2 == 0 {
			l, _ := strconv.Atoi(strNum[:len(strNum)/2])
			r, _ := strconv.Atoi(strNum[len(strNum)/2:])
			newStones = append(newStones, Stone{l, stone.Blinks + 1}, Stone{r, stone.Blinks + 1})
			continue
		}
		newStones = append(newStones, Stone{stone.Val * 2024, stone.Blinks + 1})
	}

	return newStones
}

func part1() {
	raw, _ := os.ReadFile("test.txt")
	//raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	stones := []Stone{}
	for _, char := range strings.Split(data, " ") {
		num, _ := strconv.Atoi(char)
		stones = append(stones, Stone{num, 0})
	}

	for range totalBlinks {
		stones = blink(stones)
		//fmt.Println(stones)
		//fmt.Println("***")
	}
	fmt.Println(stones)
	fmt.Println(len(stones))
	fmt.Println("**")
}

// process each number individually, then sum lengths
// once a number has an even number of digits, it will always get split down into individual digits
// ie 21 -> 2,1 (2 digits, 2 steps)
// ie 2048 -> 20,48 -> 2,0,4,8 (4 digits, 3 steps)
// ie 28676032 -> 2867,6032 -> 28,67,60,32 -> 2,8,6,7,6,0,3,2 (8 digits, 4 steps)
// ie (16 digits, 5 steps)
// non factors or 2 are more complicated

// I think keep cache of single digit number, then use knowledge of factor of 2's to skip steps
// 0 takes 4 blinks to get to 2,0,2,4
// 1,2,3,4 behave same. take 3 blinks to get to digits of n*2024 (a 4 digit number)
// 5-9 behave same, take 5 blinks to get to digits of n*2024 (a 5 digit number)

// have struct stone {
//	val int
// blinks int
//}
// I think write function that takes in an array of stones
// loop over each stone, if is single digit or power of two, use cache, and update array with new stones with value and blinks
// if blinks for stone is 75 skip
// if all stones have blinks of 75, done
// there is a way to not actually store the whole array, but my brain hurts

type Stone struct {
	Val    int
	Blinks int
}

var stoneCache = map[int][]int{
	0: []int{2, 0, 2, 4},
	1: []int{2, 0, 2, 4},
	2: []int{4, 0, 4, 8},
	3: []int{6, 0, 7, 2},
	4: []int{8, 0, 9, 6},
	5: []int{1, 0, 1, 2, 0},
	6: []int{1, 2, 1, 4, 4},
	7: []int{1, 4, 1, 6, 8},
	8: []int{1, 6, 1, 9, 2},
	9: []int{1, 8, 2, 1, 6},
}

var blinkCache = map[int]int{
	0: 4,
	1: 3,
	2: 3,
	3: 3,
	4: 3,
	5: 5,
	6: 5,
	7: 5,
	8: 5,
	9: 5,
}

func isPowOf2(n int) int {
	//added one corner case if n is zero it will also consider as power 2
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

var temp = []Stone{}

func blink2(stones []Stone, tot *int) {
	for _, stone := range stones {
		if stone.Blinks > totalBlinks {
			fmt.Println("uh oh")
		}
		if stone.Blinks == totalBlinks {
			temp = append(temp, stone)
			continue
		}
		strNum := strconv.Itoa(stone.Val)
		// use cache
		if len(strNum) == 1 {
			next := stoneCache[stone.Val]
			nextBlink := stone.Blinks + blinkCache[stone.Val]
			nextStones := []Stone{}
			// if cache would go too far, then bruteforce
			if nextBlink > totalBlinks {
				// manually step until at max blinks
				for range totalBlinks - stone.Blinks {
					nextStones = blink([]Stone{stone})
				}
			} else {
				for _, val := range next {
					nextStones = append(nextStones, Stone{val, nextBlink})
				}
			}

			*tot += len(nextStones) - 1
			blink2(nextStones, tot)
			continue
		}

		nextStones := blink([]Stone{stone})
		*tot += len(nextStones) - 1
		blink2(nextStones, tot)

		//if isPowOf2(len(strNum)) == 0 {
		//	blinks := int(math.Log2(float64(len(strNum))) + 1)
		//	digits := strings.Split(strNum, "")
		//	nextStones := []Stone{}
		//	for _, digit := range digits {
		//		// TODO: maybe have to handle 1000 -> becomes 1,0,0 not 1,0,0,0
		//		// maybe just bruteforce this part
		//		val, _ := strconv.Atoi(digit)
		//		nextStones = append(nextStones, Stone{val, stone.Blinks + blinks})
		//	}
		//	*tot += len(nextStones) - 1
		//	blink2(nextStones, tot)
		//	continue
		//}

	}
	return
}

func part2() {
	raw, _ := os.ReadFile("test.txt")
	//raw, _ := os.ReadFile("input.txt")
	data := string(raw)
	stones := []Stone{}
	for _, char := range strings.Split(data, " ") {
		num, _ := strconv.Atoi(char)
		stones = append(stones, Stone{num, 0})
	}
	//for i := range 6 {
	//	nums = blink(nums)
	//	fmt.Println(i + 1)
	//	fmt.Println(nums)
	//	fmt.Println("****")
	//}
	tot := len(stones)
	blink2(stones, &tot)
	fmt.Println(tot)
	//fmt.Println(len(nums))
}

func main() {
	part1()
	part2()
	fmt.Println(temp)
}
