package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_multiply(t *testing.T) {
	tests := []struct {
		Input    []int
		Expected int
	}{
		{[]int{1, 1, 2, 2}, 4},
		{[]int{2, 2, 1, 2}, 8},
	}

	for _, test := range tests {
		result := multiply(test.Input...)
		if result != test.Expected {
			t.Errorf("expected %d, got %d", test.Expected, result)
		}
	}
}

func Test_dirGen(t *testing.T) {
	tests := []struct {
		Width     int
		Height    int
		X         int
		Y         int
		Direction Direction
		Expected  []Coordinate
	}{
		{5, 5, 2, 2, NORTH, []Coordinate{
			{2, 1}, {2, 0},
		}},
		{5, 5, 2, 2, EAST, []Coordinate{
			{3, 2}, {4, 2},
		}},
		{5, 5, 2, 2, SOUTH, []Coordinate{
			{2, 3}, {2, 4},
		}},
		{5, 5, 2, 2, WEST, []Coordinate{
			{1, 2}, {0, 2},
		}},

		{5, 5, 0, 0, NORTH, []Coordinate{}},
		{5, 5, 4, 4, SOUTH, []Coordinate{}},
	}

	for _, test := range tests {
		results := []Coordinate{}
		for cc := range coordGen(test.Direction, test.X, test.Y, test.Width, test.Height) {
			results = append(results, cc)
		}
		if len(results) != len(test.Expected) {
			t.Fatalf("expected count of %d, got %d", len(test.Expected), len(results))
		}

		for i := range results {
			if diff := cmp.Diff(test.Expected[i], results[i]); diff != "" {
				t.Errorf("coordinates mismatch (-want, +got):\n%s", diff)
			}
		}
	}
}
