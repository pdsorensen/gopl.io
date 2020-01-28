// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("textfile.txt")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordMap := make(map[string]int)

	for scanner.Scan() {
		value := scanner.Text()
		wordMap[value]++
	}

	fmt.Print("\nword\tcount\n")
	for word, n := range wordMap {
		fmt.Printf("%-15s %d\n", word, n)
	}
}

//!-
