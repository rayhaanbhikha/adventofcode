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
		sNum := strconv.Itoa(i)
		if passedChecks(sNum) {
			passwords++
		}
	}

	fmt.Println(passwords)

}

func passedChecks(number string) bool {
	s := strings.Split(number, "")
	pV := 0

	doubleFound := false

	for _, currentVal := range s {
		cV, _ := strconv.Atoi(currentVal)

		if cV < pV {
			return false
		}

		if cV == pV {
			doubleFound = true
		}
		pV = cV
	}
	if doubleFound {
		return true
	}
	return false
}
