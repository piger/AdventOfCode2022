package main

import (
	"bufio"
	"fmt"
	"math"
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

func directionMove(d Direction) Pos {
	switch d {
	case UP:
		return directions[0]
	case RIGHT:
		return directions[2]
	case DOWN:
		return directions[4]
	case LEFT:
		return directions[6]
	}
	panic("aaah!")
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
	if p.Equal(other) {
		return true
	}

	for _, pp := range other.Surrounding() {
		if p.Equal(pp) {
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

// GetCloser returns a "step Pos" that when added to p will get closer to "other"
func (p Pos) GetCloser(other Pos) (Pos, error) {
	var candidates []Pos
	if p.X == other.X || p.Y == other.Y {
		candidates = append(candidates, straight()...)
	} else {
		candidates = append(candidates, diagonal()...)
	}

	for _, cc := range candidates {
		tmp := p.Add(cc)
		if tmp.Adjacent(other) {
			return cc, nil
		}
	}

	return Pos{0, 0}, fmt.Errorf("cannot determine how to get closer to %v", other)
}

// needToMove is a smarter way to tell if the distance between two Pos is
// greater than `amount`, which indicates that the tail needs to move.
// In the 1st problem the tail moves when the distance is 2.
// Comes from a comment on Reddit.
func needToMove(one, other Pos, amount int) bool {
	x := int(math.Abs(float64(one.X - other.X)))
	y := int(math.Abs(float64(one.Y - other.Y)))
	return x > amount || y > amount
}

func run(filename string) error {
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	head := Pos{X: 0, Y: 0}
	tail := Pos{X: 0, Y: 0}

	cells := make(map[Pos]struct{})
	cells[tail] = struct{}{}

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
		for i := 0; i < steps; i++ {
			dest := head.Add(directionMove(d))
			fmt.Printf("move head from %v to %v\n", head, dest)
			head.Set(dest)

			if !tail.Adjacent(head) {
				moveTo, err := tail.GetCloser(head)
				if err != nil {
					return err
				}
				cells[tail.Add(moveTo)] = struct{}{}
				fmt.Printf("move tail from %v to %v -- head is at %v\n", tail, tail.Add(moveTo), head)
				tail = tail.Add(moveTo)
			}
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	fmt.Printf("final positions: head=%v, tail=%v\n", head, tail)
	fmt.Printf("tail visited %d cells at least once\n", len(cells))

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
