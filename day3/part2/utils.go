package main

import (
	"math"
	"strconv"
	"strings"
)

// Coordinates ...
type Coordinates struct {
	x int
	y int
}

func getCoordinates(p string, prevCoords Coordinates) []Coordinates {
	newP := strings.SplitN(p, "", 2)
	direction := newP[0]
	increment, _ := strconv.Atoi(newP[1])

	newCoordinates := make([]Coordinates, 0)
	switch direction {
	case "U":
		// y up
		for i := 1; i <= increment; i++ {
			coord := Coordinates{prevCoords.x, i + prevCoords.y}
			newCoordinates = append(newCoordinates, coord)
		}
	case "D":
		// y down
		for i := -1; i >= -increment; i-- {
			coord := Coordinates{prevCoords.x, i + prevCoords.y}
			newCoordinates = append(newCoordinates, coord)
		}
	case "R":
		// x right
		for i := 1; i <= increment; i++ {
			coord := Coordinates{i + prevCoords.x, prevCoords.y}
			newCoordinates = append(newCoordinates, coord)
		}
	case "L":
		// x left
		for i := -1; i >= -increment; i-- {
			coord := Coordinates{i + prevCoords.x, prevCoords.y}
			newCoordinates = append(newCoordinates, coord)
		}

	}

	return newCoordinates
}

func manHattanDistance(p Coordinates) float64 {
	return math.Abs(0.0-float64(p.x)) + math.Abs(0.0-float64(p.y))
}
