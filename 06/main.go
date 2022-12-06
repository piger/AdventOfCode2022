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

func findMarker(s string) int {
	if len(s) < 4 {
		panic("invalid input")
	}

	for i := 4; i <= len(s); i++ {
		if checkSlice(s[i-4 : i]) {
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

	fmt.Printf("marker at: %d\n", findMarker(string(contents)))

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
