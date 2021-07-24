package main

//https://leetcode.com/problems/is-graph-bipartite/

func isBipartite(graph [][]int) bool {

	n := len(graph)
	visited := make([]int, n)

	// visited array protocol
	// visited index refers to the vertex id
	// visited values mean :
	//     -1 == not yet visited
	//      0 == group 0
	//      1 == group 1
	for i := 0; i < n; i++ {
		visited[i] = -1
	}

	var dfs func(vertex int, expectedGroup int) bool

	dfs = func(vertex int, expectedGroup int) bool {

		if visited[vertex] == -1 {
			visited[vertex] = expectedGroup
		} else if visited[vertex] == expectedGroup {
			// already visited
			return true
		} else {
			// not bipartite
			return false
		}

		for _, neighbor := range graph[vertex] {
			if dfs(neighbor, expectedGroup^1) == false {
				return false
			}
		}

		return true
	}

	for vertex, _ := range graph {
		if visited[vertex] == -1 {
			if dfs(vertex, 0) == false {
				return false
			}
		}
	}

	return true
}
