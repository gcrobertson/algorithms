package main

import (
	"fmt"
	"reflect"
)

// Should we use DFS, BFS?

// Node definition given by Leetcode
type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {

	// Since all nodes are connected, we are only given the first node.
	// We need to clone all connections to a new graph.
	// 1----2
	// |    |
	// 3----4
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Neighbors = []*Node{n2, n3}
	n2.Neighbors = []*Node{n1, n4}
	n3.Neighbors = []*Node{n1, n4}
	n4.Neighbors = []*Node{n2, n3}

	output := cloneGraph(n1)

	fmt.Printf("input == output? %+v\n", reflect.DeepEqual(n1, output)) // n1 == output returning false all the time.
}

func dfs(node *Node, copyMap map[int]*Node) *Node {

	// a copy of node exists in hash map, return it
	if n, ok := copyMap[node.Val]; ok {
		return n
	}

	// create node copy
	newNode := &Node{Val: node.Val}
	copyMap[node.Val] = newNode

	// here is where the recursion comes in... loop over all neighbors
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor, copyMap))
	}

	return newNode
}

// Traverse over all nodes in graph using DFS
// as we traverse, make copies of these nodes
func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)

	return dfs(node, visited)
}
