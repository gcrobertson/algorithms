package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {

	graph := make(map[int][]int, numCourses)
	// can be no longer than the number of courses
	// graph[key][k,v]

	indegree := make([]int, numCourses)

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		/*
		   this is opposite of in-degree. this is outgoing.

		   adjacency list
		   graph[5] = 2, 0
		   graph[2] = 3
		   graph[4] = 1, 0
		   graph[3] = 1
		*/
		indegree[p[0]]++
		/*
		   indegree[0] = 2 // 1 // 0
		   indegree[1] = 2 // 1
		   indegree[2] = 1 // 0
		   indegree[3] = 1
		   indegree[4] = 0
		   indegree[5] = 0
		*/
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 { // no prerequisites courses
			q = append(q, i)
		}
	}

	for len(q) > 0 {

		// q = [0,2]

		for _, c := range q {
			q = q[1:]
			numCourses--

			// c = 0
			// q = [2]

			for _, n := range graph[c] {
				indegree[n]--

				if indegree[n] == 0 {
					q = append(q, n)
				}
			}
		}
	}

	return numCourses == 0
}
