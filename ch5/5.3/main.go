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

	printTextNodes(doc)
}

func printTextNodes(node *html.Node) {
	if node == nil {
		return
	}

	if node.Type == html.ElementNode && (node.Data == "script" || node.Data == "style") {
		return
	}

	if node.Type == html.TextNode {
		fmt.Println(node.Data)
	}

	printTextNodes(node.NextSibling)
	printTextNodes(node.FirstChild)
}
