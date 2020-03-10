package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	text := "this is my text and book with $foo multiple places. $foo"

	f := func(s string) string {
		return strings.ToTitle(s)
	}

	result := expand(text, f)

	fmt.Println(result)
}
