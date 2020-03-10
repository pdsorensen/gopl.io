// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("Something went wrong when parsing the html")
	}

	counts := make(map[string]int)
	count(counts, doc)

	for element, count := range counts {
		fmt.Printf("%v : %d \n", element, count)
	}
}

func count(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		counts[n.Data]++ //n.Data is equivalant to element tag
	}

	count(counts, n.NextSibling)
	count(counts, n.FirstChild)
}
