package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	inputs := strings.Split(string(data), "\n")

	fuelPerModule := make([]float64, 0)

	for _, moduleMass := range inputs {

		m, _ := strconv.ParseFloat(moduleMass, 64)
		fuel := calculateFuel(m)
		fuelPerModule = append(fuelPerModule, fuel)
		fmt.Println(fuel)
	}
	// fmt.Println(fuelPerModule)
	fmt.Println(computeTotal(fuelPerModule))
}

func computeTotal(fuelPerModule []float64) float64 {
	sum := 0.00
	for _, fuel := range fuelPerModule {
		sum += fuel
	}
	return sum
}

func calculateFuel(mass float64) float64 {
	return math.Floor(mass/3) - 2.00
}
