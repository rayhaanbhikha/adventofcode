package main

import (
	"fmt"
)

func main() {
	program := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	amplifiers := initaliseAmplifiers([]int{9, 8, 7, 6, 5}, program)

	lastOutput := 0

	for i := 0; i < 6; i++ {
		for _, amplifier := range amplifiers {

			output := amplifier.runProgram(lastOutput)
			fmt.Println(lastOutput, "->", amplifier.name, "->", output)
			lastOutput = output
		}
	}
	fmt.Println(lastOutput)

	// amplifiers[0].runProgram(0)
	// amplifiers[0].runProgram(0)

	fmt.Println(amplifiers[0])
}

func initaliseAmplifiers(phaseSettings []int, program []int) []Amplifier {
	amplifiers := make([]Amplifier, 0)
	names := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
	}
	for i, phaseSetting := range phaseSettings {
		amplifiers = append(amplifiers, Amplifier{names[i], program, phaseSetting})
	}
	return amplifiers
}

type Amplifier struct {
	name         string
	program      []int
	phaseSetting int
}

func (a *Amplifier) runProgram(input int) int {
	return runProgram(a.program, a.phaseSetting, input)
}
