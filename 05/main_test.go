package main

import "testing"

func compareSlices(a, b []rune) bool {
	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}

func TestInsert(t *testing.T) {
	var tests = []struct {
		S        []rune
		T        rune
		Pos      int
		Expected []rune
	}{
		{[]rune{'a', 'b', 'c'}, 'r', 1, []rune{'a', 'r', 'b', 'c'}},
		{[]rune{'a', 'b', 'c'}, 'r', 0, []rune{'r', 'a', 'b', 'c'}},
	}

	for _, test := range tests {
		result := insert(test.S, test.Pos, test.T)
		if !compareSlices(result, test.Expected) {
			t.Errorf("expected %q, got %q", test.Expected, result)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		S        []rune
		Pos      int
		Expected []rune
	}{
		{[]rune{'a', 'b', 'c'}, 2, []rune{'a', 'b'}},
		{[]rune{'a', 'b', 'c'}, 0, []rune{'b', 'c'}},
	}

	for _, test := range tests {
		result := remove(test.S, test.Pos)
		if !compareSlices(result, test.Expected) {
			t.Errorf("expected %q, got %q", test.Expected, result)
		}
	}
}
