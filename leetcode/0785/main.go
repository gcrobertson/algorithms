package main

import "fmt"

func main() {

	// true
	input := [][]int{
		[]int{1, 3},
		[]int{0, 2},
		[]int{1, 3},
		[]int{0, 2},
	}

	// false
	// input := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{0, 2},
	// 	[]int{0, 1, 3},
	// 	[]int{0, 2},
	// }

	fmt.Printf("The answer is %v\n", isBipartite(input))
}

/*	The input is given as an adjacency list.
 *
 *	We will try to color a node and then color all of its neighbors a different color.
 *
 *	We will continue to do this until there is a
 *  failure because we have to re-color a node or we run out of nodes.
 *
 *	Problem includes: 	Graph Adjacency List
 *	Solution includes:	Edge coloring
 *						Breadth First Search
 */
func isBipartite(graph [][]int) bool {

	if len(graph) < 3 {
		// 0, 1, 2 will always be bipartite. is this true?
		return true
	}

	// nodes will be colored here
	//  0 	= uncolored, default
	//  1 	= blue
	// -1 	= red
	coloredNodes := make([]int, len(graph))

	// queue
	q := []int{}

	// loop through all nodes in graph
	for i := 0; i < len(coloredNodes); i++ {

		// if node has not been colored
		if coloredNodes[i] == 0 {

			// enqueue this unvisited node
			q = append(q, i)

			// color this unvisited node blue
			coloredNodes[i] = 1

			// BFS
			for len(q) > 0 {

				// dequeue first node from queue
				node := q[0]
				q = q[1:]

				// iterate through all adjacent nodes using the provided input adjacency list
				for _, neighbor := range graph[node] {

					// if the adjacent node has not been colored, then color it
					if coloredNodes[neighbor] == 0 {

						fmt.Printf("the color of the current node %v is %v\n", node, coloredNodes[node])
						coloredNodes[neighbor] = -coloredNodes[node]
						// this part is actually really clever. it will make the neighbor the opposite of the current node.
						// so 	if current node is  1,  -1 = -1
						// 		if current node is -1, --1 =  1
						q = append(q, neighbor)
						// fmt.Printf("queue is: %v\n", q)
					} else {
						// if the adjacent node has already been colored, verify the color
						if coloredNodes[neighbor] == coloredNodes[node] {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

// Below is a working DFS Solution

// func isBipartite(graph [][]int) bool {
// 	nodeNum := len(graph)
// 	visited := make([]int, nodeNum)
// 	for i := 0; i < nodeNum; i++ {
// 		if visited[i] == 0 && !explore(graph, i, 1, visited) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func explore(graph [][]int, node int, color int, visited []int) bool {
// 	if visited[node] != 0 {
// 		return visited[node] == color
// 	}
// 	visited[node] = color
// 	for i := 0; i < len(graph[node]); i++ {
// 		if !explore(graph, graph[node][i], -1*color, visited) {
// 			return false
// 		}
// 	}
// 	return true
// }
