package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func extractLinks(urli string) ([]string, error) {

	var links []string

	// issue a http get on the urli
	resp, err := http.Get(urli)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	// this is required after the http.Get function
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	// read the body into a byte array
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// stringify the byte array
	respBodyStr := string(respBody)

	index := 0 // moving pointer from start to end of the string length
	for index < len(respBodyStr) {

		// basic, custom code to extract the links from the body implemented
		// to understand string operations. The right approach would be to use the
		// html package and use functions like Parse() etc. as the below would not
		// work correctly for most modern pages which have many new ways of linking.

		// find the anchor link
		aStart := strings.Index(respBodyStr[index:], "<a ")
		if aStart == -1 {
			// did not find another Anchor tag
			break
		}
		aEnd := strings.Index(respBodyStr[index+aStart:], ">")

		// convert to absolute indexes
		aEnd = index + aStart + aEnd + 1
		aStart = index + aStart

		// extract hyperlink from the href tag
		hrefStart := strings.Index(respBodyStr[aStart:aEnd], "href=\"")
		if hrefStart != -1 {
			hrefStart = aStart + hrefStart + len("href=\"")
			hrefEnd := strings.Index(respBodyStr[hrefStart:], "\"")

			hrefEnd = hrefStart + hrefEnd
			if respBodyStr[hrefStart] == '#' {

			} else if respBodyStr[hrefStart] == '/' {
				// special casing
				u, _ := url.Parse(urli)
				links = append(links, u.Scheme+"://"+u.Host+respBodyStr[hrefStart:hrefEnd])
			} else {
				links = append(links, respBodyStr[hrefStart:hrefEnd])
			}
		}
		// move on
		index = aEnd
	}

	return links, nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Seed page missing")
		os.Exit(-1)
	}

	fmt.Println("Seed pages:", os.Args[1:])

	// link lists created and appended by the crawlers
	workQList := make(chan []string)
	// map of links seen so far
	seenLinks := make(map[string]bool)
	// unique links for the crawlers
	linksToCrawl := make(chan string)

	// needs to be a new thread, so that the main thread does not deadlock
	// sender and receiver need to different threads.
	go func() {
		workQList <- os.Args[1:]
	}()

	// create the crawlers
	for i := 0; i < 10; i++ {
		go func() {
			for link := range linksToCrawl {
				links, err := extractLinks(link)
				if err != nil {
					fmt.Println("error:", err)
				}
				go func() { workQList <- links }()
			}
		}()
	}

	// continue to dedupe the links
	for workQ := range workQList {
		for _, link := range workQ {
			if seenLinks[link] == false {
				seenLinks[link] = true
				// add to the pages to crawl stream

				fmt.Println(link)
				linksToCrawl <- link
			}
		}
	}
}
