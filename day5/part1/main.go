package main

import (
	"fmt"
)

func getInput() []int {
	// return []int{3, 0, 4, 0, 99}
	// return []int{1001, 9, 10, 3, 1002, 3, 11, 0, 99, 30, 40, 50}
	return []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 40, 93, 224, 1001, 224, -3720, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 1101, 56, 23, 225, 1102, 64, 78, 225, 1102, 14, 11, 225, 1101, 84, 27, 225, 1101, 7, 82, 224, 1001, 224, -89, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1, 35, 47, 224, 1001, 224, -140, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 1101, 75, 90, 225, 101, 9, 122, 224, 101, -72, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1102, 36, 63, 225, 1002, 192, 29, 224, 1001, 224, -1218, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 102, 31, 218, 224, 101, -2046, 224, 224, 4, 224, 102, 8, 223, 223, 101, 4, 224, 224, 1, 224, 223, 223, 1001, 43, 38, 224, 101, -52, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1102, 33, 42, 225, 2, 95, 40, 224, 101, -5850, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1102, 37, 66, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1007, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 359, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 374, 1001, 223, 1, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 389, 101, 1, 223, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 404, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 101, 1, 223, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 449, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 464, 1001, 223, 1, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 494, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 509, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 1107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 554, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 599, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 614, 101, 1, 223, 223, 107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 644, 101, 1, 223, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 659, 1001, 223, 1, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
}

func main() {
	output := runProgram(getInput(), 1)
	fmt.Println(output)
}

func runProgram(program []int, input int) int {
	p := Program{program}

	for index := 0; index < len(p.program); {
		opCode, paramModes := parseOpcode(p.get(index))

		if opCode == 99 {
			break
		}

		newIncrement := index + len(paramModes) + 1

		fmt.Println("opCode: ", opCode, "paramModes: ", paramModes)
		fmt.Println("newIncrement", newIncrement)

		if newIncrement > len(p.program) {
			break
		}

		intcode := p.program[index:newIncrement]

		computeIntcode(&p, input, opCode, paramModes, intcode)

		index = newIncrement
	}

	fmt.Println(p.program)

	return 5
}

func computeIntcode(p *Program, input, opcode int, paramModes, intcode []int) {

	getIntcode := func(intcodeIndex int) int {
		if paramModes[intcodeIndex-1] == 1 {
			return intcode[intcodeIndex]
		}
		return p.get(intcode[intcodeIndex])
	}

	switch opcode {
	case 1:
		newValue := getIntcode(1) + getIntcode(2)
		p.set(intcode[3], newValue)
		return
	case 2:
		newValue := getIntcode(1) * getIntcode(2)
		p.set(intcode[3], newValue)
		return
	case 3:
		p.set(intcode[1], input)
		return
	case 4:
		fmt.Println("ouput: ", getIntcode(1))
		return
	}
}