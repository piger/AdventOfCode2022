package main

import "testing"

func TestBasic(t *testing.T) {
	type Match struct {
		Opponent Sign
		Player   Sign
		Outcome  string
	}

	// test the outcome of each match, with the expected outcome from the opponent perspective.
	var tests = []Match{
		{Rock, Paper, "lose"},
		{Rock, Rock, "draw"},
		{Rock, Scissors, "win"},
		{Paper, Rock, "win"},
		{Paper, Paper, "draw"},
		{Paper, Scissors, "lose"},
		{Scissors, Rock, "lose"},
		{Scissors, Paper, "win"},
		{Scissors, Scissors, "draw"},
	}

	for _, test := range tests {
		outcome, _ := test.Opponent.Against(test.Player)
		if outcome != test.Outcome {
			t.Fatalf("%s against %s: expected %s, got %s", test.Opponent, test.Player, test.Outcome, outcome)
		}

		// do not test draws
		if test.Opponent != test.Player {
			wt := winTable[test.Opponent]
			lt := loseTable[test.Opponent]

			// if the outcome is "win" then player must have chosen the value from winTable, which indicates
			// what sign something wins against.
			if test.Outcome == "win" {
				if wt != test.Player {
					t.Fatalf("winTable is wrong")
				}
			} else if test.Outcome == "lose" {
				// if the outcome is "lose" then player must have chosen the value from loseTable, which indicates
				// what sign something lose against.
				if lt != test.Player {
					t.Fatalf("loseTable is wrong")
				}
			}
		}
	}
}
