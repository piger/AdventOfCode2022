package main

import (
	"fmt"
	"io"
	"os"
)

func checkSlice(chars string) bool {
	t := make(map[rune]int)

	for _, char := range chars {
		t[char] += 1
	}

	for _, val := range t {
		if val > 1 {
			return false
		}
	}
	return true
}

// findMarker finds the first sequence of characters of the specified `length` that
// are unique and returns the position of the *end* of that sequence in the input string.
//
// For example:
// bvwbjplbgvbhsrlpgdmjqwftvncz
// 123456789...
// ____^
//
// Using a for loop to iterate through the input string is convenient because the variable `i`
// always marks the position *after* the group of characters being checked.
func findMarker(s string, length int) int {
	if len(s) < length {
		panic("invalid input")
	}

	for i := length; i <= len(s); i++ {
		if checkSlice(s[i-length : i]) {
			return i
		}
	}
	return -1
}

func run() error {
	fh, err := os.Open("input")
	if err != nil {
		return err
	}
	defer fh.Close()

	contents, err := io.ReadAll(fh)
	if err != nil {
		return err
	}

	fmt.Printf("marker for length 4 is at: %d\n", findMarker(string(contents), 4))
	fmt.Printf("marker for length 14 is at: %d\n", findMarker(string(contents), 14))

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
