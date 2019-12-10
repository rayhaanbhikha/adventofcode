package main

import (
	"strconv"
	"strings"
)

func parseOpcode(opCode int) (int, []int) {
	if opCode > 99 {

		sOpCode := strconv.Itoa(opCode)
		n := len(sOpCode)
		newOpCode, _ := strconv.Atoi(sOpCode[n-2:])

		params := strings.Split(sOpCode[:n-2], "")
		pLength := len(params)
		for i := 0; i < pLength/2; i++ {
			params[i], params[pLength-(i+1)] = params[pLength-(i+1)], params[i]
		}

		paramModes := make([]int, 0)
		for _, p := range params {
			num, _ := strconv.Atoi(p)
			paramModes = append(paramModes, num)
		}

		for len(paramModes) != paramLength(newOpCode) {
			paramModes = append(paramModes, 0)
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

func paramLength(opCode int) int {
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
