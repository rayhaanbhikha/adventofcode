package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type coord struct {
	x, y float64
}

type vector struct {
	x, y float64
}

func (v vector) length() float64 {
	xSqrd := math.Pow(v.x, 2)
	ySqrd := math.Pow(v.y, 2)
	return math.Sqrt(xSqrd + ySqrd)
}

type storedVectors struct {
	vectors map[vector]coord
}

func test() {
	// mStation := coord{4, 4}
	// storedVectors := make(map[vector]coord)

	// fmt.Println(storedVectors)

	// p1 := coord{4, 0}
	// p2 := coord{4, 2}

	// v1 := computeVector(mStation, p1)
	// v2 := computeVector(mStation, p2)

	// // fmt.Println(p1, v1, v1.length())

	// // fmt.Println(p2, v2, v2.length())
	// fmt.Println(sameVectors(v1, v2))

	// addVector(mStation, p1, storedVectors)
	// addVector(mStation, p2, storedVectors)

	// // storedVectors[v1] = p1
	// fmt.Println(storedVectors)
	// addVector(origin, p2, storedVectors)
	// addVector(origin, p1, storedVectors)
	// fmt.Println(storedVectors)

	// fmt.Println(p1, v1, v1.length())
	// fmt.Println(p2, v2, v2.length())
	// fmt.Println(sameVectors(v1, v2))

}

func main() {

	data, _ := ioutil.ReadFile("../input")
	input := string(data)

	asteroidCoords := getAsteroidCoords(input)
	// checkDLS(coord{2, 2}, asteroidCoords)
	// n := len(asteroidCoords)

	maxDLS := 0
	var bestStation coord
	for _, mStation := range asteroidCoords {
		if dls := checkDLS(mStation, asteroidCoords); dls > maxDLS {
			maxDLS = dls
			bestStation = mStation
		}
	}

	fmt.Println(maxDLS, bestStation)

	fmt.Println("hello world")
}

func checkDLS(mStation coord, asteroidCoords []coord) int {
	storedVectors := make(map[vector]coord)
	for _, asteroidCoord := range asteroidCoords {
		if asteroidCoord.x == mStation.x && asteroidCoord.y == mStation.y {
		} else {
			addVector(mStation, asteroidCoord, storedVectors)
		}
	}
	return len(storedVectors)
}

func addVector(mStation, asteroidCoord coord, storedVectors map[vector]coord) {
	v := computeVector(mStation, asteroidCoord)
	// fmt.Println(mStation, asteroidCoord, v)
	found := false
	for storedV := range storedVectors {
		// fmt.Println("hi: ", sameVectors(v, storedV), v, storedV, coord)
		if sameVectors(v, storedV) {
			if storedV.length() != 0 && v.length() < storedV.length() {
				storedVectors[storedV] = asteroidCoord
			}
			found = true
		}
	}
	if !found {
		storedVectors[v] = asteroidCoord
	}
	// fmt.Println(storedVectors)
	// fmt.Print("==========\n\n")
}

func sameVectors(v1, v2 vector) bool {
	divY, divX := 0.0, 0.0

	if v1.x != 0 && v2.x != 0 {
		divX = v2.x / v1.x
	}

	if v1.y != 0 && v2.y != 0 {
		divY = v2.y / v1.y
	}

	if v1.x == 0 && v2.x == 0 {
		if v1.y > 0 && v2.y > 0 {
			return true
		}

		if v1.y < 0 && v2.y < 0 {
			return true
		}
	}

	if v1.y == 0 && v2.y == 0 {
		if v1.x > 0 && v2.x > 0 {
			return true
		}

		if v1.x < 0 && v2.x < 0 {
			return true
		}
	}

	if divX > 0 && divX == divY {
		return true
	}

	// if divY != 0 && divX == divY {
	// 	return true
	// }

	return false
}

func computeVector(p1, p2 coord) vector {
	newX := p2.x - p1.x
	newY := p2.y - p1.y

	return vector{newX, newY}
}

func getAsteroidCoords(input string) []coord {
	asteroidCoords := make([]coord, 0)
	for y, row := range strings.Split(input, "\n") {
		for x, value := range strings.Split(row, "") {
			if value == "#" {
				asteroidCoords = append(asteroidCoords, coord{x: float64(x), y: float64(y)})
			}
		}
	}
	return asteroidCoords
}
