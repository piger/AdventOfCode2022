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

// in a 10x10 grid, 0x0 is top left and 10x10 is bottom right.
var directions = []Pos{
	// north
	{0, -1},
	// north-east
	{1, -1},
	// east
	{1, 0},
	// south-east
	{1, 1},
	// south
	{0, 1},
	// south-west
	{-1, 1},
	// west
	{-1, 0},
	// north-west
	{-1, -1},
}

func diagonal() []Pos {
	var result []Pos
	for i, d := range directions {
		if i%2 != 0 {
			result = append(result, d)
		}
	}
	return result
}

func straight() []Pos {
	var result []Pos
	for i, d := range directions {
		if i%2 == 0 {
			result = append(result, d)
		}
	}
	return result
}

type Pos struct {
	X int
	Y int
}

func (p Pos) String() string {
	return fmt.Sprintf("[X: %d, Y:%d]", p.X, p.Y)
}

func (p Pos) Add(other Pos) Pos {
	var result Pos
	result.X = p.X + other.X
	result.Y = p.Y + other.Y
	return result
}

func (p *Pos) Set(dest Pos) {
	p.X = dest.X
	p.Y = dest.Y
}

func (p Pos) Equal(other Pos) bool {
	if p.X == other.X && p.Y == other.Y {
		return true
	}
	return false
}

func (p Pos) Adjacent(other Pos) bool {
	for _, pp := range p.Surrounding() {
		if other.Equal(pp) {
			return true
		}
	}
	return false
}

func (p Pos) Surrounding() []Pos {
	var result []Pos
	for _, dd := range directions {
		result = append(result, p.Add(dd))
	}
	return result
}

func run(filename string) error {
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	// head := Pos{X: 0, Y: 0}
	// tail := Pos{X: 0, Y: 0}

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
