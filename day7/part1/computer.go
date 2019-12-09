package main

func runProgram(program []int, phaseSetting, input int) int {
	getInput := initInputs(phaseSetting, input)
	p := Program{program}
	output := 0
	for index := 0; index < len(p.program); {
		opCode, paramModes := parseOpcode(p.get(index))
		newIncrement := index + instructionLength(opCode) + 1

		if opCode == 99 || newIncrement > len(p.program) {
			break
		}

		intcode := p.program[index:newIncrement]

		if newIndex, change := computeIntcode(&p, getInput, opCode, paramModes, intcode, &output); change {
			index = newIndex
		} else {
			index = newIncrement
		}

	}

	return output
}

func computeIntcode(p *Program, input func() int, opcode int, paramModes, intcode []int, output *int) (int, bool) {

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
		p.set(intcode[1], input())
		return 0, false
	case 4:
		*output = getIntcode(1)
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
