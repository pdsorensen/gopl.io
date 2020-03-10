package main

import (
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestHTMLPrettyPrint(t *testing.T) {
	resp, err := http.Get("https://golang.org")

	doc, err := html.Parse(resp.Body)
	if err != nil {
		t.Fatal("Something went wrong in test", err)
	}

	defer resp.Body.Close()

	forEachNode(doc, startElement, endElement)
}
