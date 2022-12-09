package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		Start    Pos
		Other    Pos
		Expected Pos
	}{
		{Pos{0, 0}, Pos{1, -1}, Pos{1, -1}},
		{Pos{10, 20}, Pos{-1, 1}, Pos{9, 21}},
	}

	for _, test := range tests {
		result := test.Start.Add(test.Other)
		if diff := cmp.Diff(test.Expected, result); diff != "" {
			t.Errorf("add is doing it wrong:\n%s", diff)
		}
	}
}

func TestSurrounding(t *testing.T) {
	tests := []struct {
		P        Pos
		Expected []Pos
	}{
		{Pos{10, 20}, []Pos{{10, 19}, {11, 19}, {11, 20}, {11, 21}, {11, 20}, {9, 21}, {9, 20}, {9, 19}}},
	}

	for _, test := range tests {
		result := test.P.Surrounding()
		if diff := cmp.Diff(test.Expected, result); diff != "" {
			t.Errorf("wrong surrounding:\n%s", diff)
		}
	}
}

func TestAdjacent(t *testing.T) {
	tests := []struct {
		P        Pos
		Other    Pos
		Expected bool
	}{
		{Pos{10, 20}, Pos{11, 20}, true},
	}

	for _, test := range tests {
		result := test.P.Adjacent(test.Other)
		if result != test.Expected {
			t.Errorf("adjacent is wrong: expected %v, got %v (for %+v and %+v)", test.Expected, result, test.P, test.Other)
		}
	}
}
