package main

import "testing"

func TestFindMarker(t *testing.T) {
	var tests = []struct {
		Input    string
		Expected int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, test := range tests {
		result := findMarker(test.Input)
		if result != test.Expected {
			t.Errorf("for %q got %d, expected %d", test.Input, result, test.Expected)
		}
	}
}
