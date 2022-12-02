package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputFilename = "input"
)

func run() error {
	fh, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer fh.Close()

	current := 0
	most := 0
	var counts []int

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		if line == "" {
			fmt.Printf("this elf had %d calories\n", current)
			if current > most {
				most = current
			}
			counts = append(counts, current)
			current = 0
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				return fmt.Errorf("invalid number in line %q: %w", line, err)
			}
			current += n
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner: %w", err)
	}

	fmt.Printf("the elve with most calories has %d\n", most)
	sort.Ints(counts)
	fmt.Printf("lower count: %d, upper count: %d\n", counts[0], counts[len(counts)-1])

	top3 := 0
	for i := len(counts) - 3; i < len(counts); i++ {
		fmt.Printf("top three: %d\n", counts[i])
		top3 += counts[i]
	}
	fmt.Printf("top 3: %d\n", top3)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
