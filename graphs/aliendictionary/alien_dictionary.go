package main

import (
	"fmt"
)

type Vertex struct {
	Key       rune
	Neighbors []Vertex
}

func findAlienOrder(words []string) string {

	// character(key)->Vertex(value) map
	alienChars := make(map[rune]*Vertex)

	// add all unique characters into the map
	for _, word := range words {
		for _, c := range word {
			if _, found := alienChars[c]; !found {
				alienChars[c] = &Vertex{c, make([]Vertex, 0)}
			}
		}
	}

	for w1 := 0; w1 < len(words)-1; w1++ {
		w2 := w1 + 1
		for w1idx, w2idx := 0, 0; w1idx < len(words[w1]) && w2idx < len(words[w2]); w1idx, w2idx = w1idx+1, w2idx+1 {

			if words[w1][w1idx] == words[w2][w2idx] {
				continue
			} else {
				// w2idx char comes after w1idx char
				found := false
				for _, v := range alienChars[rune(words[w1][w1idx])].Neighbors {
					if v.Key == alienChars[rune(words[w2][w2idx])].Key {
						found = true
					}
				}
				if !found {
					alienChars[rune(words[w1][w1idx])].Neighbors =
						append(alienChars[rune(words[w1][w1idx])].Neighbors, *alienChars[rune(words[w2][w2idx])])
					fmt.Println(string(words[w1][w1idx]), string(words[w2][w2idx]))
				}
				break
			}
		}
	}

	// build the relationship between the nodes
	for k, v := range alienChars {
		fmt.Println(string(k), v)
	}

	// do a topological sort

	//
	return ""
}

func main() {
	testData := []struct {
		input []string
		want  string
	}{
		{[]string{"baa", "abcd", "abca", "cab", "cad"}, "bdac"},
		{[]string{"caa", "aaa", "aab"}, "cab"},
	}

	for _, test := range testData {
		result := findAlienOrder(test.input)
		if result != test.want {
			fmt.Println("Test Failed, expected:", test.want, "got:", result)
		} else {
			fmt.Println("Test passed")
		}
	}
}
