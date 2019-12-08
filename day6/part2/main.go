package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var totalOrbits int = 0

type Orbit struct {
	orbitObject    string
	orbitingObject string
}

func main() {

	rootNode := NewNode("", "COM")

	data, _ := ioutil.ReadFile("../input.txt")
	orbits := strings.Split(string(data), "\n")

	orbitArrayMap := make([]Orbit, 0)

	for _, orbit := range orbits {
		orbitData := strings.Split(orbit, ")")
		orbitArrayMap = append(orbitArrayMap, Orbit{orbitData[0], orbitData[1]})
	}

	iterateMap(rootNode, orbitArrayMap)

	youNode := rootNode.find("YOU", "")

	findSAN(youNode, 0)

	fmt.Println(totalOrbits)
}

func findSAN(currentNode *Node, orbits int) {
	parentNode := currentNode.parentNode
	fmt.Println(parentNode.data, parentNode.level)

	foundNode := parentNode.find("SAN", currentNode.data)
	if foundNode == nil {
		findSAN(parentNode, orbits+1)
	} else {
		res := foundNode.level - 1 - parentNode.level + orbits
		fmt.Println("Number of orbits, ", res)
		return
	}

}

func iterateMap(rootNode *Node, orbitMap []Orbit) {
	if len(orbitMap) == 0 {
		return
	}

	notFound := make([]Orbit, 0)
	for _, orbitData := range orbitMap {
		if !rootNode.contains(orbitData.orbitObject) {
			notFound = append(notFound, orbitData)
			continue
		}
		rootNode.insert(NewNode(orbitData.orbitObject, orbitData.orbitingObject))
	}

	if len(notFound) == 0 {
		return
	}

	// fmt.Println(notFound)
	iterateMap(rootNode, notFound)
}
