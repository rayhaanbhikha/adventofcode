package main

import (
	"strconv"
	"strings"
)

func initInputs(inputs ...int) func() int {
	called := 0
	return func() int {
		if called > len(inputs)-1 {
			return 0
		}
		defer func() { called++ }()
		return inputs[called]
	}
}

func parseOpcode(opCode int) (int, []int) {
	if opCode > 99 {

		sOpCode := strconv.Itoa(opCode)
		n := len(sOpCode)
		parsedOpCode := sOpCode[n-2:]
		newOpCode, _ := strconv.Atoi(parsedOpCode)

		params := strings.Split(sOpCode[:n-2], "")
		paramModes := make([]int, 0)
		for i := len(params) - 1; i >= 0; i-- {
			mode, _ := strconv.Atoi(params[i])
			paramModes = append(paramModes, mode)
		}

		return newOpCode, paramModes
	}

	switch opCode {
	case 1, 2, 7, 8:
		return opCode, []int{0, 0, 0}
	case 3, 4:
		return opCode, []int{0}
	case 5, 6:
		return opCode, []int{0, 0}
	case 99:
		fallthrough
	default:
		return opCode, []int{}
	}
}

func instructionLength(opCode int) int {
	switch opCode {
	case 1, 2, 7, 8:
		return 3
	case 3, 4:
		return 1
	case 5, 6:
		return 2
	case 99:
		fallthrough
	default:
		return 0
	}
}
