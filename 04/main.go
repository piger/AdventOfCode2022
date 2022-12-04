package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeSet struct {
	Start int
	End   int
}

func NewRangeSet(s string) (*RangeSet, error) {
	fields := strings.Split(s, "-")
	if len(fields) != 2 {
		return nil, fmt.Errorf("wrong number of elements in RangeSet: %d", len(fields))
	}
	start, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}

	r := &RangeSet{
		Start: start,
		End:   end,
	}
	return r, nil
}

func (r *RangeSet) Contains(other *RangeSet) bool {
	if r.Start >= other.Start && r.End <= other.End {
		return true
	}
	return false
}

func (r *RangeSet) ContainsAny(other *RangeSet) bool {
	if r.Start <= other.End && r.End >= other.End {
		return true
	}
	return false
}

func run() error {
	fh, err := os.Open("input")
	if err != nil {
		return err
	}
	defer fh.Close()

	var counter int
	var counter2 int

	s := bufio.NewScanner(fh)
	for s.Scan() {
		line := s.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			return fmt.Errorf("line has wrong number of fields: %d (%q)", len(fields), line)
		}

		pair1, err := NewRangeSet(fields[0])
		if err != nil {
			return err
		}

		pair2, err := NewRangeSet(fields[1])
		if err != nil {
			return err
		}

		if pair1.Contains(pair2) || pair2.Contains(pair1) {
			counter++
		}
		if pair1.ContainsAny(pair2) || pair2.ContainsAny(pair1) {
			counter2++
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	// 582 assignments have ranges that fully contain each other
	// 893 assignments have ranges that overlaps at all
	fmt.Printf("%d assignments have ranges that fully contain each other\n", counter)
	fmt.Printf("%d assignments have ranges that overlaps at all\n", counter2)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
