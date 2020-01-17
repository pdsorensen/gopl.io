// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.
// Rev reverses a slice.
package main

func main() {
	//!+array
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// b := [...]int{0, 1, 2, 3, 4, 5}
	// reverse(a[:])
	// reverse2(&b)
	// fmt.Println(a) // "[5 4 3 2 1 0]"
	// fmt.Println(b) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	// fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	// 	input := bufio.NewScanner(os.Stdin)
	// outer:
	// 	for input.Scan() {
	// 		var ints []int
	// 		for _, s := range strings.Fields(input.Text()) {
	// 			x, err := strconv.ParseInt(s, 10, 64)
	// 			if err != nil {
	// 				fmt.Fprintln(os.Stderr, err)
	// 				continue outer
	// 			}
	// 			ints = append(ints, int(x))
	// 		}
	// 		reverse(ints)
	// 		fmt.Printf("%v\n", ints)
	// 	}
	// NOTE: ignoring potential errors from input.Err()

	// 4.5
	// strings := []string{"a", "b", "b", "a", "a", "b"}
	// unique := removeAdjecant(strings)
	// println(unique, strings)

	// 4.6
	strings := []string{"Hello world", "space spac e space"}
	squashed := squashSpaces(strings)
	println(squashed, strings)
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int) {
	first := s[0]
	copy(s, s[1:])
	s[len(s)-1] = first

	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }
}

func removeAdjecant(strings []string) []string {
	index := 0

	for _, s := range strings {
		if strings[index] == s {
			continue
		}

		index++
		strings[index] = s
	}

	return strings[:index+1]
}

//!-rev
