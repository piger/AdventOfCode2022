package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// isVisible determine if a given tree at point x,y is visible from the edge of the grid;
// a tree is visible if all of the other trees between it and an edge of the grid are shorter than it.
func isVisible(x, y, width, height int, grid [][]int) bool {
	var (
		visibleN bool
		visibleE bool
		visibleS bool
		visibleW bool
	)

	tree := grid[y][x]
	fmt.Printf("analysing tree at %dx%d (%d)\n", x, y, tree)

	// north
	for yPos := y - 1; yPos >= 0; yPos-- {
		if grid[yPos][x] >= tree {
			fmt.Printf("not visible from N: %dx%d (%d)\n", x, yPos, grid[yPos][x])
			visibleN = false
			break
		} else {
			visibleN = true
			fmt.Printf("visible from N: %dx%d (%d)\n", x, yPos, grid[yPos][x])
		}
	}

	// east
	for xPos := x + 1; xPos < width; xPos++ {
		if grid[y][xPos] >= tree {
			fmt.Printf("not visible from E: %dx%d (%d)\n", xPos, y, grid[y][xPos])
			visibleE = false
			break
		} else {
			visibleE = true
			fmt.Printf("visible from E: %dx%d (%d)\n", xPos, y, grid[y][xPos])
		}
	}

	// south
	for yPos := y + 1; yPos < height; yPos++ {
		if grid[yPos][x] >= tree {
			fmt.Printf("not visible from S: %dx%d (%d)\n", x, yPos, grid[yPos][x])
			visibleS = false
			break
		} else {
			visibleS = true
			fmt.Printf("visible from S: %dx%d (%d)\n", x, yPos, grid[yPos][x])
		}
	}

	// west
	for xPos := x - 1; xPos >= 0; xPos-- {
		if grid[y][xPos] >= tree {
			fmt.Printf("not visible from W: %dx%d (%d)\n", xPos, y, grid[y][xPos])
			visibleW = false
			break
		} else {
			visibleW = true
			fmt.Printf("visible from W: %dx%d (%d)\n", xPos, y, grid[y][xPos])
		}
	}

	visible := visibleN || visibleE || visibleS || visibleW
	fmt.Printf("tree at %dx%d (%d) is visible? %v\n", x, y, tree, visible)

	return visible
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

			// north
			if y > 0 {
				for yy := y - 1; yy >= 0; yy-- {
					count++
					if grid[yy][x] >= grid[y][x] {
						// scores = append(scores, count)
						// count = 0
						break
					}
				}
				if count > 0 {
					scores = append(scores, count)
					count = 0
				}
			}

			// east
			if x < width-1 {
				for xx := x + 1; xx < width; xx++ {
					count++
					if grid[y][xx] >= grid[y][x] {
						break
					}
				}
				if count > 0 {
					scores = append(scores, count)
					count = 0
				}
			}

			// south
			if y < height-1 {
				for yy := y + 1; yy < height; yy++ {
					count++
					if grid[yy][x] >= grid[y][x] {
						break
					}
				}
				if count > 0 {
					scores = append(scores, count)
					count = 0
				}
			}

			// west
			if x > 0 {
				for xx := x - 1; xx >= 0; xx-- {
					count++
					if grid[y][xx] >= grid[y][x] {
						break
					}
				}
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
