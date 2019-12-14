package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type coord struct {
	x, y      float64
	destroyed bool
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

type pointInfo struct {
	v       vector
	vLength float64
	vAngle  float64
	c       coord
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
	n := len(asteroidCoords)
	fmt.Println(n)
	laserCoord := coord{x: 19, y: 11}

	fmt.Println(asteroidCoords)
	startingIndex := 1

	for !allAsteroidsDestroyed(asteroidCoords) {
		collectedVectors := getAsteroidVectors(laserCoord, asteroidCoords)
		targets := targetCollectedAsteroids(collectedVectors)
		startingIndex = destroyTargets(startingIndex, targets, asteroidCoords)
	}

	fmt.Println(asteroidCoords)

	// for _, p := range points {
	// 	fmt.Println("vec: ", p.v, "\tcoord: ", p.c, "\tangle: ", p.vAngle)
	// }

	// for _, asteroidCoord := range asteroidCoords {
	// 	fmt.Println(asteroidCoord)
	// }

	fmt.Println("hello world")
}

func allAsteroidsDestroyed(asteroidCoords []coord) bool {
	for _, asteroidCoord := range asteroidCoords {
		if !asteroidCoord.destroyed {
			return false
		}
	}
	return true
}

func destroyTargets(startingIndex int, targets []pointInfo, asteroidCoords []coord) int {
	count := startingIndex
	for vaporisedCount, p := range targets {
		for index, asteroidCoord := range asteroidCoords {
			if p.c.x == asteroidCoord.x && p.c.y == asteroidCoord.y && !asteroidCoord.destroyed {
				asteroidCoords[index].destroyed = true
				fmt.Println(vaporisedCount+startingIndex, " ==>> ", asteroidCoord)
				count++
				// updateAsteroid(asteroidCoords, index)
			}
		}
	}
	return count
}

func targetCollectedAsteroids(collectedVectors map[vector]coord) []pointInfo {
	points := make([]pointInfo, 0)
	for v, c := range collectedVectors {
		angle := 0.00

		switch {
		case v.x > 0 && v.y < 0:
			angle = math.Atan(math.Abs(v.x)/math.Abs(v.y)) * 180 / math.Pi
		case v.x > 0 && v.y > 0:
			angle = math.Atan(math.Abs(v.y)/math.Abs(v.x))*180/math.Pi + 90
		case v.x <= 0 && v.y >= 0:
			angle = math.Atan(math.Abs(v.x)/math.Abs(v.x))*180/math.Pi + 180
		case v.x < 0 && v.y < 0:
			angle = math.Atan(math.Abs(v.y)/math.Abs(v.x))*180/math.Pi + 270
		}

		if angle == 0 && v.x > 0 {
			angle = 90
		}

		if angle == 0 && v.x < 0 {
			angle = 270
		}

		points = append(points, pointInfo{v, v.length(), angle, c})
	}

	sort.Sort(byAngle(points))
	return points
}

type byAngle []pointInfo

func (p byAngle) Len() int {
	return len(p)
}

func (p byAngle) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p byAngle) Less(i, j int) bool {
	// if p[i].vAngle < p[j].vAngle {
	// 	return true
	// }

	// // if p[i].vAngle == p[j].vAngle {
	// // 	return p[i].vLength < p[j].vLength
	// // }

	return p[i].vAngle < p[j].vAngle
}

func getAsteroidVectors(mStation coord, asteroidCoords []coord) map[vector]coord {
	storedVectors := make(map[vector]coord)
	for _, asteroidCoord := range asteroidCoords {
		if asteroidCoord.x == mStation.x && asteroidCoord.y == mStation.y {
		} else if !asteroidCoord.destroyed {
			addVector(mStation, asteroidCoord, storedVectors)
		}
	}
	return storedVectors
}

func addVector(mStation, asteroidCoord coord, storedVectors map[vector]coord) {
	v := computeVector(mStation, asteroidCoord)
	// fmt.Println(mStation, asteroidCoord, v)
	found := false
	for storedV := range storedVectors {
		// fmt.Println("hi: ", sameVectors(v, storedV), v, storedV, coord)
		if sameVectors(v, storedV) {
			if storedV.length() != 0 && v.length() < storedV.length() {
				delete(storedVectors, storedV)
				storedVectors[v] = asteroidCoord
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
				asteroidCoords = append(asteroidCoords, coord{x: float64(x), y: float64(y), destroyed: false})
			}
		}
	}
	return asteroidCoords
}
