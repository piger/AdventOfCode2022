package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isVisible(x, y, width, height int, grid [][]int) bool {
	var visible bool

	tree := grid[y][x]
	fmt.Printf("analysing tree at %dx%d (%d)\n", x, y, tree)

	// north
	for yPos := y - 1; yPos >= 0; yPos-- {
		if grid[yPos][x] > tree {
			fmt.Printf("not visible from N: %dx%d (%d)\n", x, yPos, grid[yPos][x])
			visible = false
			break
		} else {
			visible = true
			fmt.Printf("visible from N: %dx%d (%d)\n", x, yPos, grid[yPos][x])
		}
	}

	// east
	for xPos := x + 1; xPos < width; xPos++ {
		if grid[y][xPos] > tree {
			fmt.Printf("not visible from E: %dx%d (%d)\n", xPos, y, grid[y][xPos])
			visible = false
			break
		} else {
			visible = true
			fmt.Printf("visible from E: %dx%d (%d)\n", xPos, y, grid[y][xPos])
		}
	}

	// south
	for yPos := y + 1; yPos < height; yPos++ {
		if grid[yPos][x] > tree {
			fmt.Printf("not visible from S: %dx%d (%d)\n", x, yPos, grid[yPos][x])
			visible = false
			break
		} else {
			visible = true
			fmt.Printf("visible from S: %dx%d (%d)\n", x, yPos, grid[yPos][x])
		}
	}

	// west
	for xPos := x - 1; xPos >= 0; xPos-- {
		if grid[y][xPos] > tree {
			fmt.Printf("not visible from W: %dx%d (%d)\n", xPos, y, grid[y][xPos])
			visible = false
			break
		} else {
			visible = true
			fmt.Printf("visible from W: %dx%d (%d)\n", xPos, y, grid[y][xPos])
		}
	}

	fmt.Printf("tree at %dx%d (%d) is visible? %v\n", x, y, tree, visible)

	return visible
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

	// actual problem
	var visible int
	// trees on the edges are always visible
	visible += (width * 2) + (height-2)*2
	fmt.Printf("visible trees on the edges: %d\n", visible)

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if !isVisible(x, y, width, height, grid) {
				visible++
			}
		}
		fmt.Println()
	}

	fmt.Printf("total visible trees: %d\n", visible)

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
