package main

import "fmt"

type Node struct {
	parentNodeName string
	parentNode     *Node
	data           string
	children       []*Node
	level          int
}

func (n *Node) insert(newNode *Node) {
	if newNode.parentNodeName == n.data {
		newNode.level += n.level
		newNode.parentNode = n
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
	fmt.Printf("[%s -> %s] level: %d\n", n.parentNodeName, n.data, n.level)
	if len(n.children) > 0 {
		fmt.Println("New level")
		for _, childNode := range n.children {
			childNode.print()
		}
	}
}

func (n Node) find(nodeName string, exclude string) *Node {
	if nodeName == n.data {
		return &n
	} else if len(n.children) > 0 {
		var foundNode *Node
		for _, childNode := range n.children {
			if childNode.data == exclude {
				continue
			}
			if x := childNode.find(nodeName, exclude); x != nil {
				foundNode = x
				break
			}
		}
		return foundNode
	} else {
		return nil
	}
}

func NewNode(parentNodeName, name string) *Node {
	return &Node{parentNodeName, nil, name, make([]*Node, 0), 1}
}
