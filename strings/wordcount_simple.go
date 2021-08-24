package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordmap := make(map[string]int)

	for i := range words {
		fmt.Println(words[i])
		if _, ok := wordmap[words[i]]; ok {
			wordmap[words[i]] += 1
		} else {
			wordmap[words[i]] = 1
		}
	}
	return wordmap
}

func main() {
	wc.Test(WordCount)
}
