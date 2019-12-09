package main

import (
	"strconv"
	"strings"
)

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

		if n < 5 && newOpCode < 3 {
			for i := n; i < 5; i++ {
				paramModes = append(paramModes, 0)
			}
		}

		return newOpCode, paramModes
	}

	switch opCode {
	case 1, 2:
		return opCode, []int{0, 0, 0}
	case 3, 4:
		return opCode, []int{0}
	case 99:
		fallthrough
	default:
		return opCode, []int{}
	}
}
