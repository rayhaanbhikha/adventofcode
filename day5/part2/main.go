package main

import (
	"fmt"
)

func main() {
	program := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 40, 93, 224, 1001, 224, -3720, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 56, 23, 225, 1102, 64, 78, 225, 1102, 14, 11, 225, 1101, 84, 27, 225, 1101, 7, 82, 224, 1001, 224, -89, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1, 35, 47, 224, 1001, 224, -140, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 1101, 75, 90, 225, 101, 9, 122, 224, 101, -72, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1102, 36, 63, 225, 1002, 192, 29, 224, 1001, 224, -1218, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 102, 31, 218, 224, 101, -2046, 224, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1001, 43, 38, 224, 101, -52, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1102, 33, 42, 225, 2, 95, 40, 224, 101, -5850, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1102, 37, 66, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 359, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 374, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 389, 101, 1, 223, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 404, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 449, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 464, 1001, 223, 1, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 509, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 554, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 599, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 659, 1001, 223, 1, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
	// program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	runProgram(program, 5)
}

func runProgram(program []int, input int) int {
	p := Program{program}

	for index := 0; index < len(p.program); {
		opCode, paramModes := parseOpcode(p.get(index))
		newIncrement := index + instructionLength(opCode) + 1

		if opCode == 99 || newIncrement > len(p.program) {
			break
		}

		fmt.Println("opCode: ", opCode, "paramModes: ", paramModes)
		fmt.Println("newIncrement", newIncrement)

		intcode := p.program[index:newIncrement]

		if newIndex, change := computeIntcode(&p, input, opCode, paramModes, intcode); change {
			fmt.Println("change")
			index = newIndex
		} else {
			index = newIncrement
		}

	}

	fmt.Println(p.program)

	return 5
}

func computeIntcode(p *Program, input, opcode int, paramModes, intcode []int) (int, bool) {

	getIntcode := func(intcodeIndex int) int {
		if intcodeIndex-1 < len(paramModes) && paramModes[intcodeIndex-1] == 1 {
			return intcode[intcodeIndex]
		}
		return p.get(intcode[intcodeIndex])
	}

	switch opcode {
	case 1:
		newValue := getIntcode(1) + getIntcode(2)
		p.set(intcode[3], newValue)
		return 0, false
	case 2:
		newValue := getIntcode(1) * getIntcode(2)
		p.set(intcode[3], newValue)
		return 0, false
	case 3:
		p.set(intcode[1], input)
		return 0, false
	case 4:
		fmt.Println("ouput: ", getIntcode(1))
		return 0, false
	case 5:
		firstParam := getIntcode(1)
		secondParam := getIntcode(2)
		if firstParam > 0 {
			return secondParam, true
		}
		return 0, false
	case 6:
		firstParam := getIntcode(1)
		secondParam := getIntcode(2)
		if firstParam == 0 {
			return secondParam, true
		}
		return 0, false
	case 7:
		firstParam := getIntcode(1)
		secondParam := getIntcode(2)
		if firstParam < secondParam {
			p.set(intcode[3], 1)
		} else {
			p.set(intcode[3], 0)
		}
		return 0, false
	case 8:
		firstParam := getIntcode(1)
		secondParam := getIntcode(2)
		if firstParam == secondParam {
			p.set(intcode[3], 1)
		} else {
			p.set(intcode[3], 0)
		}
		return 0, false
	default:
		return 0, false
	}
}
