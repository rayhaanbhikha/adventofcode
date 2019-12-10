package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type layer struct {
	digits []int
}

func main() {
	fmt.Println("hello world")

	data, _ := ioutil.ReadFile("../input.txt")
	input := string(data)
	w, h := 25, 6

	// input := "" +
	// 	"122121" +
	// 	"022221" +
	// 	"112220" +
	// 	"221221" +
	// 	"000012"

	// w, h := 3, 2

	separator := w * h
	layers := make([]layer, 0)
	parsedInput := strings.Split(input, "")
	n := len(parsedInput)

	for i := 0; i < n; {
		if i+separator > n {
			break
		}
		newlayer := NewLayer(parsedInput[i : i+separator])
		layers = append(layers, newlayer)
		i += separator
	}

	pixels := layers[0].digits

	for i := 1; i < len(layers); i++ {
		layer := layers[i]
		for index, digit := range layer.digits {
			currentPixel := pixels[index]
			if currentPixel == 2 && digit != 2 {
				pixels[index] = digit
			}
		}
	}

	for i := 0; i < len(pixels); {
		if i+w > len(pixels) {
			break
		}
		fmt.Println(pixels[i:i+w], len(pixels[i:i+w]))
		i += w
	}

}

func NewLayer(numbers []string) layer {
	digits := make([]int, 0)
	for _, num := range numbers {
		convertedNum, _ := strconv.Atoi(num)
		digits = append(digits, convertedNum)
	}

	return layer{digits}
}
