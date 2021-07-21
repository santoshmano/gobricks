package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"discrete math(dm)":         {"math 1(m1)", "intro to programming(itp)"},
	"data structures(ds)":       {"discrete math(dm)"},
	"algorithms(algo)":          {"data structures(ds)", "discrete math(dm)"},
	"programming languages(pl)": {"data structures(ds)", "computer organization(corg)"},
	"operating systems(os)":     {"data structures(ds)", "computer organization(corg)"},
}

func main() {

	visited := make(map[string]bool)
	var sortedCourses []string

	for vertex := range prereqs {
		visited[vertex] = false
	}
	for vertex := range visited {
		topoSort(vertex, visited, &sortedCourses)
	}

	fmt.Println("--------------")
	fmt.Println("Sorted Courses")
	fmt.Println("--------------")
	for _, course := range sortedCourses {
		fmt.Println(course)
	}
}

func topoSort(vertex string, visited map[string]bool, sortedCourses *[]string) {

	if visited[vertex] == true {
		return
	}

	visited[vertex] = true

	for _, course := range prereqs[vertex] {
		_, present := visited[course]
		if !present {
			visited[vertex] = false
		}

		if visited[vertex] == false {
			topoSort(course, visited, sortedCourses)
		}
	}

	*sortedCourses = append(*sortedCourses, vertex)
}
