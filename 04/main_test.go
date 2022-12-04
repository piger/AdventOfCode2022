package main

import "testing"

func TestNewRangeSet(t *testing.T) {
	var tests = []struct {
		Input string
		Start int
		End   int
	}{
		{"10-20", 10, 20},
		{"13-29", 13, 29},
	}

	for _, test := range tests {
		result, err := NewRangeSet(test.Input)
		if err != nil {
			t.Fatal(err)
		}

		if result.Start != test.Start || result.End != test.End {
			t.Fatalf("expected %d-%d, got %d-%d", test.Start, test.End, result.Start, result.End)
		}
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		Input1   string
		Input2   string
		Expected bool
	}{
		{"2-8", "3-7", true},
		{"3-7", "2-8", false},
		{"6-6", "4-6", false},
		{"10-20", "30-40", false},
	}

	for _, test := range tests {
		one, err := NewRangeSet(test.Input1)
		if err != nil {
			t.Fatal(err)
		}

		other, err := NewRangeSet(test.Input2)
		if err != nil {
			t.Fatal(err)
		}

		if one.Contains(other) != test.Expected {
			t.Errorf("%v in %v: expected %v", one, other, test.Expected)
		}
	}
}

func TestContainsAny(t *testing.T) {
	var tests = []struct {
		Input1   string
		Input2   string
		Expected bool
	}{
		{"5-7", "7-9", true},
	}

	for _, test := range tests {
		one, err := NewRangeSet(test.Input1)
		if err != nil {
			t.Fatal(err)
		}

		other, err := NewRangeSet(test.Input2)
		if err != nil {
			t.Fatal(err)
		}

		if one.ContainsAny(other) != test.Expected && other.ContainsAny(one) != test.Expected {
			t.Errorf("%v in %v: expected %v", one, other, test.Expected)
		}
	}
}
