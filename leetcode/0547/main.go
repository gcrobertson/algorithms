package main

import "fmt"

func main() {

	input := [][]int{
		[]int{1, 1, 0}, // (0,0) self. 0 -> 1 direct friendship
		[]int{1, 1, 0}, // (1,1) self. 1 -> 0 has been visited already.
		[]int{0, 0, 1}, // (2,2) self.
	}

	// visit = []bool{
	//	friend 0 = true, @ i = 0, this becomes true. dfs(matrix, 0, this array)
	//	friend 1 = false,
	//	friend 2 = false,
	//}

	/*
		input := [][]int{
			[]int{1, 0, 0, 1},	// (0,0) self. (0,3) are friends.
			[]int{0, 1, 1, 0},	// (1,1) self. (1,2) are friends.
			[]int{0, 1, 1, 1},      // (2,2) self. (2,1) are friends. (2,3) are friends.
			[]int{1, 0, 1, 1},      // (3,3) self. (3,0), (3,2) are friends.
		}
	*/

	// 0 -> 3 -> 2 -> 1

	result := findCircleNum(input)

	fmt.Printf("The result is %v\n", result)
}

func findCircleNum(M [][]int) int {

	var ans int
	visit := make([]bool, len(M))
	for i := 0; i < len(M); i++ {
		if !visit[i] {
			visit[i] = true
			dfs(M, i, visit)
			//bfs(M, i, visit)
			ans++
		}
	}
	return ans
}

func dfs(M [][]int, me int, visit []bool) {
	for i := 0; i < len(M); i++ {
		if !visit[i] && M[me][i] == 1 {
			visit[i] = true
			dfs(M, i, visit)
		}
	}
}

func bfs(M [][]int, me int, visit []bool) {
	queue := []int{me}
	for len(queue) > 0 {
		dequeue := queue[0]
		queue = queue[1:]
		visit[dequeue] = true
		for i := 0; i < len(M); i++ {
			if !visit[i] && M[dequeue][i] == 1 {
				visit[i] = true
				queue = append(queue, i)
			}
		}
		if len(queue) == 0 {
			return
		}
	}
}
