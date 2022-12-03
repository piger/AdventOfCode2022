package main

import "testing"

func TestScores(t *testing.T) {
	tests := map[rune]int{
		'a': 1,
		'z': 26,
		'A': 27,
		'Z': 52,
	}

	for r, val := range tests {
		if scores[r] != val {
			t.Fatalf("expected score for %c: %d, got: %d", r, val, scores[r])
		}
	}
}

func compareRunes(one, other []rune) bool {
	if len(one) != len(other) {
		return false
	}

	stamp := make(map[rune]bool)
	for _, r := range one {
		stamp[r] = true
	}

	for _, r := range other {
		if _, ok := stamp[r]; !ok {
			return false
		}
	}

	return true
}

func TestCommonRunes(t *testing.T) {
	type test struct {
		one    string
		other  string
		result []rune
	}

	tests := []test{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp", []rune{'p'}},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL", []rune{'L'}},
		{"PmmdzqPrV", "vPwwTWBwg", []rune{'P'}},
	}

	for _, thing := range tests {
		result := commonRunes(thing.one, thing.other)
		if !compareRunes(result, thing.result) {
			t.Fatalf("expected %v, got %v", thing.result, result)
		}
	}
}
