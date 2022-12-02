package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Sign int

const (
	Rock Sign = iota
	Paper
	Scissors

	Lose = 0
	Draw = 3
	Win  = 6
)

func (s Sign) String() string {
	return [...]string{"Rock", "Paper", "Scissors"}[s]
}

func (s Sign) Against(other Sign) (string, int) {
	switch s {
	case Rock:
		switch other {
		case Rock:
			return "draw", Draw
		case Paper:
			return "lose", Lose
		case Scissors:
			return "win", Win
		}

	case Paper:
		switch other {
		case Rock:
			return "win", Win
		case Paper:
			return "draw", Draw
		case Scissors:
			return "lose", Lose
		}

	case Scissors:
		switch other {
		case Rock:
			return "lose", Lose
		case Paper:
			return "win", Win
		case Scissors:
			return "draw", Draw
		}
	}
	panic("boh!")
}

// maps each sign to what it wins against
var winTable = map[Sign]Sign{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

// maps each sign to what it loses against
var loseTable = map[Sign]Sign{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var scoreTable = map[Sign]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var opponentHand = map[string]Sign{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var playerHand = map[string]Sign{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func run() error {
	fh, err := os.Open("input")
	if err != nil {
		return err
	}
	defer fh.Close()

	var score int
	s := bufio.NewScanner(fh)
	for s.Scan() {
		values := strings.Split(s.Text(), " ")
		if len(values) != 2 {
			return fmt.Errorf("wrong number of fields: %d", len(values))
		}

		opp, ok := opponentHand[values[0]]
		if !ok {
			panic("baddo value")
		}

		pl, ok := playerHand[values[1]]
		if !ok {
			panic("baddo player value")
		}

		desc, v := opp.Against(pl)
		fmt.Printf("%s against %s: %s (%d)\n", opp, pl, desc, v)

		fmt.Printf("Opponent: %s - Player: %s => ", opp, pl)

		if opp == pl {
			score += Draw
			score += scoreTable[opp]
			fmt.Println("draw")
		} else {
			winAgainst := winTable[opp]
			fmt.Printf(" [would win against: %s] ", winAgainst)

			if pl == winAgainst {
				// we lost
				score += Lose
				fmt.Println("player lose")
			} else {
				score += Win
				fmt.Println("player win")
			}

			score += scoreTable[pl]
		}

		fmt.Println("")
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Printf("score: %d\n", score)

	return nil
}

func run2() error {
	fh, err := os.Open("input")
	if err != nil {
		return err
	}
	defer fh.Close()

	var score int
	s := bufio.NewScanner(fh)
	for s.Scan() {
		line := s.Text()
		values := strings.Split(line, " ")
		if len(values) != 2 {
			return fmt.Errorf("bad line: %q", line)
		}

		opp := opponentHand[values[0]]
		// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
		switch values[1] {
		case "X":
			// lose
			loseAgainst := winTable[opp]
			score += Lose
			score += scoreTable[loseAgainst]
		case "Y":
			// draw
			score += Draw
			score += scoreTable[opp]
		case "Z":
			// win
			winAgainst := loseTable[opp]
			score += Win
			score += scoreTable[winAgainst]
		}
	}

	fmt.Printf("score2: %d\n", score)
	if score == 11866 || score == 10497 {
		fmt.Println("nah")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	if err := run2(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
