package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	ID     int
	Crates []rune
}

func longestLine(lines []string) (longest int) {
	for _, line := range lines {
		l := len(line)
		if l > longest {
			longest = l
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

	// firsy we read the initial stack configuration
	var buf []string

	s := bufio.NewScanner(fh)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		buf = append(buf, line)
	}

	numStacks := (longestLine(buf) + 1) / 4
	stacks := make([]Stack, numStacks)

	// assign IDs to each stack
	for i := range stacks {
		stacks[i].ID = i + 1
	}

	// for each line
	for _, line := range buf {
		// for each crate column
		for i, j := 0, 0; i <= len(line); i, j = i+4, j+1 {
			if line[i] == ' ' {
				continue
			}
			// fmt.Printf("crate %d: %c\n", j, line[i+1])
			stacks[j].Crates = append(stacks[j].Crates, rune(line[i+1]))
		}
	}

	for _, stack := range stacks {
		fmt.Printf("stack %d: %q\n", stack.ID, stack.Crates)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
