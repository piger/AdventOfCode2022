package main

import "testing"

func TestGetBoundaries(t *testing.T) {
	type test struct {
		Input string
		Start int
		End   int
	}

	tests := []test{
		{"10-20", 10, 20},
	}

	for _, thing := range tests {
		start, end, err := getBoundaries(thing.Input)
		if err != nil {
			t.Fatal(err)
		}

		if start != thing.Start {
			t.Fatalf("expected %d, got %d", thing.Start, start)
		}

		if end != thing.End {
			t.Fatalf("expected %d, got %d", thing.End, end)
		}
	}
}

func TestIntersection(t *testing.T) {
	type test struct {
		Small    map[int]bool
		Large    map[int]bool
		Expected bool
	}

	tests := []test{
		{
			map[int]bool{10: true, 11: true, 12: true, 13: true},
			map[int]bool{9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true},
			true,
		},
		{
			map[int]bool{10: true, 11: true, 12: true, 13: true},
			map[int]bool{11: true, 12: true, 13: true, 14: true, 15: true},
			false,
		},
	}

	for _, thing := range tests {
		result := intersection(thing.Small, thing.Large)
		if result != thing.Expected {
			t.Fatalf("expected %v, got %v", thing.Expected, result)
		}
	}
}
