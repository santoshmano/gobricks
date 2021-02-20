package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordmap := make(map[string]int)

	for i := range words {
		fmt.Println(words[i])
		if _, ok := wordmap[words[i]]; ok==true {
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
