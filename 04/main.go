package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getBoundaries(s string) (start, end int, err error) {
	fields := strings.Split(s, "-")
	if len(fields) != 2 {
		err = fmt.Errorf("wrong number of fields in range: %d (%q)", len(fields), s)
		return
	}

	start, err = strconv.Atoi(fields[0])
	if err != nil {
		return
	}

	end, err = strconv.Atoi(fields[1])
	if err != nil {
		return
	}

	return
}

func intersection(small, large map[int]bool) bool {
	for key := range small {
		if _, ok := large[key]; !ok {
			return false
		}
	}
	return true
}

func anyIntersection(one, other map[int]bool) bool {
	for key := range one {
		if _, ok := other[key]; ok {
			return true
		}
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

		start1, end1, err := getBoundaries(fields[0])
		if err != nil {
			return err
		}
		slot1 := make(map[int]bool)
		for i := start1; i <= end1; i++ {
			slot1[i] = true
		}
		len1 := end1 - start1

		start2, end2, err := getBoundaries(fields[1])
		if err != nil {
			return err
		}
		slot2 := make(map[int]bool)
		for i := start2; i <= end2; i++ {
			slot2[i] = true
		}
		len2 := end2 - start2

		if len1 > len2 {
			contained := intersection(slot2, slot1)
			if contained {
				counter++
			}
		} else {
			contained := intersection(slot1, slot2)
			if contained {
				counter++
			}
		}

		if anyIntersection(slot1, slot2) {
			counter2++
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

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
