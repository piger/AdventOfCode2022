package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (d Direction) String() string {
	return [...]string{"NORTH", "EAST", "SOUTH", "WEST"}[d]
}

type Coordinate struct {
	X int
	Y int
}

func dirGen(d Direction, x, y, width, height int) chan Coordinate {
	out := make(chan Coordinate)

	go func() {
		switch d {
		case NORTH:
			if y == 0 {
				break
			}

			for yy := y - 1; yy >= 0; yy-- {
				out <- Coordinate{X: x, Y: yy}
			}

		case EAST:
			if x >= width {
				break
			}

			for xx := x + 1; xx < width; xx++ {
				out <- Coordinate{X: xx, Y: y}
			}

		case SOUTH:
			if y >= height {
				break
			}

			for yy := y + 1; yy < height; yy++ {
				out <- Coordinate{X: x, Y: yy}
			}

		case WEST:
			if x == 0 {
				break
			}

			for xx := x - 1; xx >= 0; xx-- {
				out <- Coordinate{X: xx, Y: y}
			}
		default:
			panic("unknown direction")
		}

		close(out)
	}()

	return out
}

func drain[T any](ch chan T) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}

// isVisible determine if a given tree at point x,y is visible from the edge of the grid;
// a tree is visible if all of the other trees between it and an edge of the grid are shorter than it.
func isVisible(x, y, width, height int, grid [][]int) bool {
	tree := grid[y][x]
	fmt.Printf("analysing tree at %dx%d (%d)\n", x, y, tree)
	visibility := make(map[Direction]bool)

Loop:
	for _, d := range []Direction{NORTH, EAST, SOUTH, WEST} {
		fmt.Printf("Checking direction: %s\n", d)

		coords := dirGen(d, x, y, width, height)
		for cc := range coords {
			fmt.Printf("checking %dx%d with %dx%d: %v\n", x, y, cc.X, cc.Y, grid[cc.Y][cc.X] > tree)
			// if the tree is NOT visible in this direction:
			if grid[cc.Y][cc.X] >= tree {
				drain(coords)
				visibility[d] = false
				continue Loop
			}
		}
		visibility[d] = true
	}

	var result bool
	for k, v := range visibility {
		if v {
			fmt.Printf("visible at %s\n", k)
			result = true
		}
	}

	return result
}

func multiply(nums ...int) (result int) {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	result = nums[0]

	for i := 1; i < len(nums); i++ {
		result *= nums[i]
	}
	return result
}

func findHighestScore(grid [][]int) int {
	var result int

	width, height := len(grid[0]), len(grid)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			count := 0
			scores := []int{}

			for _, d := range []Direction{NORTH, EAST, SOUTH, WEST} {
				coords := dirGen(d, x, y, width, height)
				for cc := range coords {
					count++
					if grid[cc.Y][cc.X] >= grid[y][x] {
						break
					}
				}
				drain(coords)
				if count > 0 {
					scores = append(scores, count)
					count = 0
				}
			}

			score := multiply(scores...)
			if score > result {
				result = score
			}

		}
	}

	return result
}

func run(filename string) error {
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	var lines []string

	// read the whole thing
	s := bufio.NewScanner(fh)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if err := s.Err(); err != nil {
		return err
	}

	// calculate size
	width, height := len(lines[0]), len(lines)
	fmt.Printf("map size: %dx%d\n", width, height)

	// allocate grid
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// populate grid
	for y, line := range lines {
		for x, char := range line {
			tree, err := strconv.Atoi(string(char))
			if err != nil {
				return err
			}
			grid[y][x] = tree
		}
	}

	/*
		// print grid
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				fmt.Printf("%d", grid[y][x])
			}
			fmt.Println()
		}
	*/

	// actual problem
	var visible int
	// trees on the edges are always visible
	visible += (width * 2) + (height-2)*2
	fmt.Printf("visible trees on the edges: %d\n", visible)

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if isVisible(x, y, width, height, grid) {
				visible++
			}
		}
		fmt.Println()
	}

	fmt.Printf("total visible trees: %d\n", visible)
	fmt.Printf("highest scenic score: %d\n", findHighestScore(grid))

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
