package main

import (
	"fmt"
)

type Params struct {
	value int
	mode  int
}

type Intcode struct {
	opCode     int
	parameters []Params
}

type Program struct {
	input        []int
	inputIndex   int
	programEnd   bool
	program      []int
	currentIndex int
	intcode      Intcode
	output       int
}

func NewProgram(program []int, input ...int) Program {
	return Program{programEnd: false, program: program, input: input, inputIndex: 0, currentIndex: 0}
}

func (p *Program) nextIntcode() {
	opCode, paramModes := parseOpcode(p.get(p.currentIndex))
	intcodeLength := p.currentIndex + len(paramModes) + 1

	if intcodeLength >= len(p.program) {
		p.increaseMemory(intcodeLength)
	}

	if opCode == 99 {
		p.programEnd = true
		return
	}

	intcode := p.program[p.currentIndex:intcodeLength]
	parameters := make([]Params, 0)

	for i, param := range intcode[1:] {
		parameters = append(parameters, Params{param, paramModes[i]})
	}

	p.intcode = Intcode{opCode, parameters}

	p.currentIndex = intcodeLength

}

func (p Program) getParamValue(index int) int {

	param := p.intcode.parameters[index]

	if param.mode == 1 {
		return param.value
	}
	return p.get(param.value)
}

func (p Program) getParam(index int) int {

	param := p.intcode.parameters[index]
	return param.value
}

func (p *Program) getInput() int {
	fmt.Println(p.inputIndex)
	value := p.input[p.inputIndex]
	p.inputIndex = 1
	return value
}

func (p *Program) computeIntcode() {

	switch p.intcode.opCode {
	case 1:
		newValue := p.getParamValue(0) + p.getParamValue(1)
		p.set(p.getParam(2), newValue)
	case 2:
		newValue := p.getParamValue(0) * p.getParamValue(1)
		p.set(p.getParam(2), newValue)
	case 3:
		p.set(p.getParam(0), p.getInput())
	case 4:
		p.output = p.getParamValue(0)
	case 5:
		firstParam, secondParam := p.getParamValue(0), p.getParamValue(1)
		if firstParam > 0 {
			p.currentIndex = secondParam
		}
	case 6:
		firstParam, secondParam := p.getParamValue(0), p.getParamValue(1)
		if firstParam == 0 {
			p.currentIndex = secondParam
		}
	case 7:
		firstParam, secondParam := p.getParamValue(0), p.getParamValue(1)
		if firstParam < secondParam {
			p.set(p.getParam(2), 1)
		} else {
			p.set(p.getParam(2), 0)
		}
	case 8:
		firstParam, secondParam := p.getParamValue(0), p.getParamValue(1)
		if firstParam == secondParam {
			p.set(p.getParam(2), 1)
		} else {
			p.set(p.getParam(2), 0)
		}
	case 99:
		p.programEnd = true
	}
}

func (p *Program) increaseMemory(index int) {
	zeroArray := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for len(p.program)-1 <= index {
		p.program = append(p.program, zeroArray...)
	}
	return
}

func (p Program) get(index int) int {
	if index < 0 {
		fmt.Println("negative index: ", index)
	}
	if index > len(p.program)-1 {
		p.increaseMemory(index)
		return p.program[index]
	}
	return p.program[index]
}

func (p *Program) set(index, value int) {
	if index > len(p.program)-1 {
		p.increaseMemory(index)
	}
	p.program[index] = value
}
