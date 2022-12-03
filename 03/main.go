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

func unique(s string) string {
	t := make(map[rune]bool)
	var result []rune

	for _, char := range s {
		t[char] = true
	}

	for key, val := range t {
		if val {
			result = append(result, key)
		}
	}
	return string(result)
}

func commonRunesN(sets []string) (result []rune) {
	t := make(map[rune]int)

	for _, set := range sets {
		set = unique(set)
		for _, char := range set {
			t[char] += 1
		}
	}

	for char, count := range t {
		if count >= len(sets) {
			result = append(result, char)
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

	var score int
	var score2 int
	var counter int
	var buf []string

	s := bufio.NewScanner(fh)
	for s.Scan() {
		line := s.Text()
		l := len(line)
		comp1 := line[0:(l / 2)]
		comp2 := line[l/2 : l]

		/*
			fmt.Printf("%s (len: %d)\n", line, l)
			fmt.Printf("comp1: %s\n", comp1)
			fmt.Printf("comp2: %s\n", comp2)
		*/

		common := commonRunes(comp1, comp2)
		for _, r := range common {
			// fmt.Printf("%c", r)
			priority, ok := scores[r]
			if !ok {
				return fmt.Errorf("missing score for %c", r)
			}

			score += priority
		}
		// fmt.Println()

		// part two
		buf = append(buf, line)
		counter++
		if counter >= 3 {
			fmt.Printf("Finding common runes in %v\n", buf)
			common := commonRunesN(buf)
			for _, char := range common {
				sc, ok := scores[char]
				if !ok {
					return fmt.Errorf("missing score for %c", char)
				}
				score2 += sc
			}

			// reset
			counter = 0
			buf = nil
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Printf("score: %d, score2: %d\n", score, score2)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
