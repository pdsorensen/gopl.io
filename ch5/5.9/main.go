package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f(s), -1)
}
