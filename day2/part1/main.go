package main

import (
	"fmt"
)

type Program struct {
	program []int
}

func (p Program) get(index int) int {
	if index > len(p.program) {
		fmt.Println("TOO high")
	}
	return p.program[index]
}

func (p *Program) set(index, value int) {
	if index < len(p.program) {
		p.program[index] = value
	}
}

func getInput() []int {
	return []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 19, 5, 23, 1, 23, 9, 27, 2, 27, 6, 31, 1, 31, 6, 35, 2, 35, 9, 39, 1, 6, 39, 43, 2, 10, 43, 47, 1, 47, 9, 51, 1, 51, 6, 55, 1, 55, 6, 59, 2, 59, 10, 63, 1, 6, 63, 67, 2, 6, 67, 71, 1, 71, 5, 75, 2, 13, 75, 79, 1, 10, 79, 83, 1, 5, 83, 87, 2, 87, 10, 91, 1, 5, 91, 95, 2, 95, 6, 99, 1, 99, 6, 103, 2, 103, 6, 107, 2, 107, 9, 111, 1, 111, 5, 115, 1, 115, 6, 119, 2, 6, 119, 123, 1, 5, 123, 127, 1, 127, 13, 131, 1, 2, 131, 135, 1, 135, 10, 0, 99, 2, 14, 0, 0}
}

func main() {
	output := runProgram(getInput(), 12, 2)
	fmt.Println(output)
}

func runProgram(program []int, noun, verb int) int {
	p := Program{program}
	p.set(1, noun)
	p.set(2, verb)
	i := 0
	for {

		if len(p.program) < i+4 {
			break
		}

		intcode := p.program[i : i+4]
		opCode := intcode[0]
		newValue := 0

		if opCode == 99 {
			break
		}

		if opCode == 2 {
			newValue = p.get(intcode[1]) * p.get(intcode[2])
		}

		if opCode == 1 {
			newValue = p.get(intcode[1]) + p.get(intcode[2])
		}

		p.set(intcode[3], newValue)

		if i < len(p.program) {
			i += 4
		}
	}

	return p.program[0]
}
