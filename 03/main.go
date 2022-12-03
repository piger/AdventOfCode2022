package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var scores map[rune]int

func init() {
	scores = make(map[rune]int)

	var j int = 1
	for i := 'a'; i <= 'z'; i++ {
		scores[i] = j
		j++
	}

	j = 27
	for i := 'A'; i <= 'Z'; i++ {
		scores[i] = j
		j++
	}
}

func commonRunes(one, other string) (result []rune) {
	t := make(map[rune]int)

	for _, r := range one {
		t[r] = 1
	}

	for _, r := range other {
		if _, ok := t[r]; ok {
			t[r] += 1
		}
	}

	for r, val := range t {
		if val > 1 {
			result = append(result, r)
		}
	}

	return
}

func run() error {
	fh, err := os.Open("input")
	if err != nil {
		return err
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)
	for s.Scan() {
		line := s.Text()
		l := len(line)
		comp1 := line[0:(l / 2)]
		comp2 := line[l/2 : l]

		fmt.Printf("%s (len: %d)\n", line, l)
		fmt.Printf("comp1: %s\n", comp1)
		fmt.Printf("comp2: %s\n", comp2)

		ruck := make(map[rune]int)
		for _, r := range comp1 {
			ruck[r] = 1
		}

		for _, r := range comp2 {
			if _, ok := ruck[r]; ok {
				ruck[r]++
			}
		}

		for r, val := range ruck {
			if val > 1 {
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
