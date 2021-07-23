package main

// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/

func countComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	visited := make([]bool, n)

	// Build the graph
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// Call DFS
	connectedComponents := 0
	for v, _ := range graph {
		if visited[v] == false {
			dfs(v, &graph, &visited)
			connectedComponents += 1
		}
	}
	return connectedComponents
}

func dfs(vertex int, graph *[][]int, visited *[]bool) {

	(*visited)[vertex] = true

	for _, neighbor := range (*graph)[vertex] {
		if (*visited)[neighbor] == false {
			dfs(neighbor, graph, visited)
		}
	}
}

/* with anonymous function, no need of pointers, will avoid anonymous
functions for now, till i am comfortable with golang pointers

import "fmt"

func countComponents(n int, edges [][]int) int {
    graph := make([][]int, n)
    visited := make([]bool, n)

    // Build the graph
    for _, e := range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        graph[e[1]] = append(graph[e[1]], e[0])
    }

    var dfs func(int)

    dfs = func(vertex int) {
        visited[vertex] = true
        for _, neighbor := range graph[vertex] {
            if visited[neighbor] == false {
                dfs(neighbor)
            }
        }
    }

    connectedComponents := 0
    for v, _ := range graph {
        if visited[v] == false {
            dfs(v)
            connectedComponents += 1
        }
    }
    return connectedComponents
}

*/
