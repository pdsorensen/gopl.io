// Echo1 prints its command-line arguments.
package main

import "fmt"

//!+bench

import (
	"testing"
	"strings"
)

//!-bench

func BenchmarkEcho1(b *testing.B) {
	args := []string {"testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3"}

	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	
	fmt.Println(args[0])
	fmt.Println(s)
}

func BenchmarkEcho2(b *testing.B) {
	args := []string {"testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3"}
	s, sep := "", ""
	for index, arg := range args[1:] {
		fmt.Print(index)
		fmt.Println(":" + sep + arg)
	}
	fmt.Println(s)
}

func BenchmarkEcho3(b *testing.B) {
	args := []string {"testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3", "testing", "arg1", "arg2", "arg3"}
	fmt.Println(strings.Join(args[1:], " "))
}

// echo1: 0.000007
// echo2: 0.000042
// echo3: 0.002s