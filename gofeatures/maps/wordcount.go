package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// this wordcount test needs to split only based on space
// dont bother with ToLower and Punctuations, commas etc.
func WordCount(s string) map[string]int {
	m := make(map[string]int)

	strings := strings.Split(s, " ")

	for _, str := range strings {
		_, ok := m[str]
		if ok {
			m[str] += 1
		} else {
			m[str] = 1
		}
	}

	return m

	/* sample test output
	/private/var/folders/21/8903byb903gfp4_6skhlk1bw0000gn/T/___go_build_wordcount_go
	PASS
	 f("I am learning Go!") =
	  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
	PASS
	 f("The quick brown fox jumped over the lazy dog.") =
	  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
	PASS
	 f("I ate a donut. Then I ate another donut.") =
	  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
	PASS
	 f("A man a plan a canal panama.") =
	  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}

	Process finished with exit code 0
	*/

}

func main() {
	wc.Test(WordCount)
}
