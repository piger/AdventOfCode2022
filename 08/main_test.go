package main

import "testing"

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
