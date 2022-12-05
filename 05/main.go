package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	modeFlag = flag.Int("mode", 1, "Run mode: 1 for the 1st problem and 2 for the 2nd")
)

var moveCmdRe = regexp.MustCompile(`^move (?P<num>\d+) from (?P<src>\d+) to (?P<dst>\d+)`)

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

func insert(a []rune, index int, value rune) []rune {
	if len(a) == index {
		return append(a, value)
	}

	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func remove(a []rune, index int) []rune {
	return append(a[:index], a[index+1:]...)
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

	// should probably check s.Err() here?

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

	/*
		for _, stack := range stacks {
			fmt.Printf("stack %d: %q\n", stack.ID, stack.Crates)
		}
	*/

	// read the move commands
	for s.Scan() {
		line := s.Text()
		// fmt.Println(line)
		matches := moveCmdRe.FindStringSubmatch(line)
		idxNum := moveCmdRe.SubexpIndex("num")
		idxSrc := moveCmdRe.SubexpIndex("src")
		idxDst := moveCmdRe.SubexpIndex("dst")
		// fmt.Printf("move %s items from %s to %s\n", matches[idxNum], matches[idxSrc], matches[idxDst])
		num, err := strconv.Atoi(matches[idxNum])
		if err != nil {
			return err
		}

		src, err := strconv.Atoi(matches[idxSrc])
		if err != nil {
			return err
		}

		dst, err := strconv.Atoi(matches[idxDst])
		if err != nil {
			return err
		}

		if *modeFlag == 1 {
			// move each crate
			for i := 0; i < num; i++ {
				// fmt.Printf("before: %q %q\n", stacks[src-1].Crates, stacks[dst-1].Crates)
				crate := stacks[src-1].Crates[0]
				stacks[dst-1].Crates = insert(stacks[dst-1].Crates, 0, crate)
				stacks[src-1].Crates = remove(stacks[src-1].Crates, 0)
				// fmt.Printf("after: %q %q\n", stacks[src-1].Crates, stacks[dst-1].Crates)
			}
		} else {
			for i := num; i > 0; i-- {
				// fmt.Printf("before: %q %q\n", stacks[src-1].Crates, stacks[dst-1].Crates)
				crate := stacks[src-1].Crates[i-1]
				stacks[dst-1].Crates = insert(stacks[dst-1].Crates, 0, crate)
				stacks[src-1].Crates = remove(stacks[src-1].Crates, i-1)
				// fmt.Printf("after: %q %q\n", stacks[src-1].Crates, stacks[dst-1].Crates)
			}
		}
	}

	var solution []rune
	for i := range stacks {
		solution = append(solution, stacks[i].Crates[0])
	}
	fmt.Printf("solution (mode %d): %s\n", *modeFlag, string(solution))

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	if *modeFlag != 1 && *modeFlag != 2 {
		fmt.Println("invalid mode flag: the allowed values are '1' and '2'")
		os.Exit(1)
	}

	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
