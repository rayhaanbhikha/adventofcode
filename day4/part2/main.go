package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")
	lowerLimit := 193651
	upperLimit := 649729

	passwords := 0

	for i := lowerLimit + 1; i < upperLimit-1; i++ {
		if passedChecks(i) {
			passwords++
		}
	}

	fmt.Println(passwords)
}

func validateNumbersFound(numbersFound map[int]int) bool {

	largerSetFound := false
	doubleFound := false

	for _, repeated := range numbersFound {

		if repeated == 2 {
			doubleFound = true
		}

		if repeated > 2 {
			largerSetFound = true
		}
	}

	switch {
	case !doubleFound:
		return false
	case !largerSetFound && doubleFound:
		return true
	case largerSetFound && doubleFound:
		return true
	default:
		return false
	}
}

func passedChecks(number int) bool {
	sNum := strconv.Itoa(number)
	s := strings.Split(sNum, "")

	numbersFound := make(map[int]int)

	pV := 0
	for _, currentVal := range s {
		cV, _ := strconv.Atoi(currentVal)

		if cV < pV {
			return false
		}

		numbersFound[cV]++

		pV = cV
	}

	return validateNumbersFound(numbersFound)
}
