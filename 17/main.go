package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Reg struct {
	A int
	B int
	C int
	// instruction pointer
	IP int
}

type Op func(op int, reg *Reg)

var instructions = map[int]Op{
	0: adv,
	1: bxl,
	2: bst,
	3: jnz,
	4: bxc,
	5: out,
	6: bdv,
	7: cdv,
}

func comboOp(op int, reg Reg) int {
	switch {
	case op <= 3:
		return op
	case op == 4:
		return reg.A
	case op == 5:
		return reg.B
	case op == 6:
		return reg.C
	}

	return 0
}

func adv(op int, reg *Reg) {
	reg.A = reg.A / (int(math.Pow(2, float64(comboOp(op, *reg)))))
	reg.IP += 2
}

func bxl(op int, reg *Reg) {
	reg.B = reg.B ^ op
	reg.IP += 2
}

func bst(op int, reg *Reg) {
	reg.B = comboOp(op, *reg) % 8
	reg.IP += 2
}

func jnz(op int, reg *Reg) {
	if reg.A == 0 {
		reg.IP += 2
		return
	}
	reg.IP = op
}

func bxc(op int, reg *Reg) {
	reg.B = reg.B ^ reg.C
	reg.IP += 2
}

func out(op int, reg *Reg) {
	fmt.Print(comboOp(op, *reg) % 8)
	fmt.Print(",")
	reg.IP += 2
}

func bdv(op int, reg *Reg) {
	reg.B = reg.A / (int(math.Pow(2, float64(comboOp(op, *reg)))))
	reg.IP += 2
}

func cdv(op int, reg *Reg) {
	reg.C = reg.A / (int(math.Pow(2, float64(comboOp(op, *reg)))))
	reg.IP += 2
}

func part1() {
	//raw, _ := os.ReadFile("test.txt")
	raw, _ := os.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n\n")

	regs := strings.Split(data[0], "\n")
	a, _ := strconv.Atoi(strings.TrimSpace(strings.Split(regs[0], ":")[1]))
	b, _ := strconv.Atoi(strings.TrimSpace(strings.Split(regs[1], ":")[1]))
	c, _ := strconv.Atoi(strings.TrimSpace(strings.Split(regs[2], ":")[1]))
	reg := Reg{A: a, B: b, C: c}

	prog := strings.Split(strings.Split(data[1], "Program: ")[1], ",")
	program := make([]int, len(prog))
	for i, part := range prog {
		program[i], _ = strconv.Atoi(strings.TrimSpace(part))
	}

	for reg.IP < len(program) {
		opCode := program[reg.IP]
		operand := program[reg.IP+1]
		instruction := instructions[opCode]
		instruction(operand, &reg)
	}
}

func main() {
	part1()
}
