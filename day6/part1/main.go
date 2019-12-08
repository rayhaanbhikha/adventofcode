package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var totalOrbits int = 0

type Node struct {
	orbitingObjectName string
	data               string
	children           []*Node
	level              int
}

func (n *Node) insert(newNode *Node) {
	if newNode.orbitingObjectName == n.data {
		newNode.level += n.level
		totalOrbits += newNode.level - 1
		n.children = append(n.children, newNode)
		return
	}

	if len(n.children) > 0 {
		for _, childNode := range n.children {
			childNode.insert(newNode)
		}
		return
	}
}

func (n *Node) contains(nodeName string) bool {
	if n.data == nodeName {
		return true
	} else if len(n.children) > 0 {
		containsNode := false
		for _, childNode := range n.children {
			if childNode.contains(nodeName) {
				containsNode = true
				break
			}
		}
		return containsNode
	} else {
		return false
	}
}

func (n Node) print() {
	fmt.Printf("[%s -> %s] level: %d\n", n.orbitingObjectName, n.data, n.level)
	if len(n.children) > 0 {
		fmt.Println("New level")
		for _, childNode := range n.children {
			childNode.print()
		}
	}
}

func NewNode(orbitingObjectName, name string) *Node {
	return &Node{orbitingObjectName, name, make([]*Node, 0), 1}
}

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

	rootNode.print()
	fmt.Println(totalOrbits)
	fmt.Println("hello world")
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

	fmt.Println(notFound)
	iterateMap(rootNode, notFound)
}
