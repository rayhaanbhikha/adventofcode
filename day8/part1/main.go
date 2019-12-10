package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type layer struct {
	digits       []int
	numOf0digits int
	numOf1digits int
	numOf2digits int
}

func main() {
	fmt.Println("hello world")

	// input := "123456789012"
	data, _ := ioutil.ReadFile("../input.txt")

	input := string(data)

	layers := make([]layer, 0)
	parsedInput := strings.Split(input, "")
	n := len(parsedInput)

	separator := 25 * 6

	for i := 0; i < n; {
		if i+separator > n {
			break
		}
		newlayer := NewLayer(parsedInput[i : i+separator])
		layers = append(layers, newlayer)
		i += separator
	}

	minZeroes := 0
	layerFound := layer{}

	for _, layer := range layers {
		fmt.Println(layer.numOf0digits)
		if minZeroes == 0 || layer.numOf0digits < minZeroes {
			minZeroes = layer.numOf0digits
			layerFound = layer
		}
	}

	// fmt.Println(layers)
	fmt.Println(minZeroes)
	fmt.Println(layerFound)

}

func NewLayer(numbers []string) layer {
	digits := make([]int, 0)
	numOf0digits := 0
	numOf1digits := 0
	numOf2digits := 0
	for _, num := range numbers {
		convertedNum, _ := strconv.Atoi(num)

		switch convertedNum {
		case 0:
			numOf0digits++
		case 1:
			numOf1digits++
		case 2:
			numOf2digits++
		}
		digits = append(digits, convertedNum)
	}

	return layer{digits, numOf0digits, numOf1digits, numOf2digits}
}
