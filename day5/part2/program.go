package main

import "fmt"

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
