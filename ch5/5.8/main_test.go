package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestFindElementByID(t *testing.T) {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	expected := "main"

	result := findElementByID(doc, "page")
	if result.Data != expected {
		t.Errorf("%s did not equal expected %s", result.Data, expected)
	}

	fmt.Printf("result: %s \n", result.Data)

}
