package main

import (
	"math"
	"strings"
)

func main() {

}

func alienOrder(words []string) string {

	// Step 0. Initialize data structures & find all unique letters
	adjList := make(map[string][]string)
	counts := make(map[string]int)

	for _, word := range words {
		for _, c := range word {
			counts[string(c)] = 0
			adjList[string(c)] = []string{}
		}
	}

	//fmt.Printf("counts initialized: %v\n", counts)   // map[e:0 f:0 r:0 t:0 w:0]
	//fmt.Printf("adjList initialized: %v\n", adjList) // map[e:[] f:[] r:[] t:[] w:[]]

	// Step 1. Find all Edges
	for i := 0; i < len(words)-1; i++ {

		word1 := words[i]
		word2 := words[i+1]

		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}

		min := int(math.Min(float64(len(word1)), float64(len(word2))))

		for j := 0; j < min; j++ {

			if word1[j] != word2[j] {

				// put into adjacency list
				// add to count
				adjList[string(word1[j])] = append(adjList[string(word1[j])], string(word2[j]))
				counts[string(word2[j])]++

				break
			}

		}

	}

	//fmt.Printf("counts populated: %v\n", counts)   // counts populated: map[e:1 f:1 r:1 t:1 w:0]
	//fmt.Printf("adjList populated: %v\n", adjList) // adjList populated: map[e:[r] f:[] r:[t] t:[f] w:[e]]

	// Adjacency List:

	// e > r
	// f
	// r > t
	// t > f
	// w > e

	q := []string{}

	for letter, count := range counts {
		if count == 0 {
			q = append(q, letter)
		}
	}
	// fmt.Printf("initial q: %v\n", q) // [w]

	var output string

	for len(q) > 0 {

		c := q[0]
		q = q[1:]

		output += c

		for i := 0; i < len(adjList[c]); i++ {

			next := adjList[c][i]
			counts[next]--

			if counts[next] == 0 {
				q = append(q, next)
			}
		}
	}

	if len(output) < len(counts) {
		return ""
	}

	return output
}

// words = []string
// wrt          // "wer"
// wrf          // "tf"
// er
// ett          // "rt"
// rftt
// wertf

// Topological sort

// Overview of Approaches

// - Adjacency Lists
// - Graph coloring
//
// https://www.geeksforgeeks.org/graph-and-its-representations/
// https://www.geeksforgeeks.org/topological-sorting-indegree-based-solution/

// `Graph`: A graph is a data structure that consists of the following two components:
// 1. A finite set of vertices also called as nodes.
// 2. A finite set of ordered pair of the form (u, v) called as edge. The pair is ordered because (u, v) is not the same as (v, u) in case of a directed graph(di-graph). The pair of the form (u, v) indicates that there is an edge from vertex u to vertex v. The edges may contain weight/value/cost.

// The following two are the most commonly used representations of a graph.
// 1. Adjacency Matrix
// 2. Adjacency List
