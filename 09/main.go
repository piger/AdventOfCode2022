package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (d Direction) String() string {
	return [...]string{"UP", "DOWN", "LEFT", "RIGHT"}[d]
}

func run(filename string) error {
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)
	for s.Scan() {
		fields := strings.Split(s.Text(), " ")
		if len(fields) != 2 {
			return fmt.Errorf("wrong number of fields in line: %q", fields)
		}

		var d Direction
		switch fields[0] {
		case "U":
			d = UP
		case "D":
			d = DOWN
		case "L":
			d = LEFT
		case "R":
			d = RIGHT
		default:
			return fmt.Errorf("unknown direction: %s", fields[0])
		}

		steps, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		fmt.Printf("%v -> %d\n", d, steps)
	}

	return nil
}

func main() {
	filename := "input"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if err := run(filename); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
