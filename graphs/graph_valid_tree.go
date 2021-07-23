package main

// https://leetcode.com/problems/graph-valid-tree/

func validTree(n int, edges [][]int) bool {
	// Graph is a valid tree if it has 0 or more vertices,
	// all connected and no cycles

	if n == 0 {
		return true
	}

	// array of integer slices, of size n
	// index represents the vertex and stored list of neighbor vertices.
	graph := make([][]int, n)

	// array with vertices as index and value sas true/false == visited/unvisited
	visited := make([]bool, n)

	// build the graph
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// traverse using dfs/bfs
	var dfs func(int, int) bool

	dfs = func(vertex, parent int) bool {
		visited[vertex] = true

		for _, neighbor := range graph[vertex] {
			if neighbor == parent {
				continue
			} else if visited[neighbor] == true {
				// there is a cycle
				return false
			} else {
				if dfs(neighbor, vertex) == false {
					return false
				}
			}
		}
		return true
	}

	// if there are any unvisited vertices return false
	if dfs(0, -1) == false {
		return false
	}

	for _, v := range visited {
		if v == false {
			return false
		}
	}

	return true
}
