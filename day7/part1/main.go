package main

import (
	"fmt"
)

func main() {
	program := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 30, 51, 72, 81, 94, 175, 256, 337, 418, 99999, 3, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 1002, 9, 2, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 101, 4, 9, 9, 102, 5, 9, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}

	maxThrust := 0
	var maxThrustSequence []int

	for _, sequence := range sequences {
		amplifiers := initaliseAmplifiers(sequence)

		lastOutput := 0

		for _, amplifier := range amplifiers {
			output := amplifier.runProgram(program, lastOutput)

			lastOutput = output
		}
		if lastOutput > maxThrust {
			maxThrustSequence = sequence
			maxThrust = lastOutput
		}
	}

	// fmt.Println(amplifiers)
	fmt.Println(maxThrustSequence, maxThrust)

}

func initaliseAmplifiers(phaseSettings []int) []Amplifier {
	amplifiers := make([]Amplifier, 0)
	for _, phaseSetting := range phaseSettings {
		amplifiers = append(amplifiers, Amplifier{phaseSetting})
	}
	return amplifiers
}

type Amplifier struct {
	phaseSetting int
}

func (a Amplifier) runProgram(program []int, input int) int {
	p := NewProgram(program, a.phaseSetting, input)

	fmt.Println(len(p.program))
	for !p.programEnd {
		p.nextIntcode()
		p.computeIntcode()
	}
	fmt.Println(len(p.program))

	return p.output
}
