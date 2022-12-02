/* rock paper scissors

1st column:
 A = rock
 B = paper
 C = scissors

2nd column:
 X = rock
 Y = paper
 Z = scissors

Scores:
- 1 rock
- 2 paper
- 3 scissors

- 0 lost
- 3 draw
- 6 won

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.

* ROCK ROCK
- ROCK PAPER
- ROCK SCISSORS

- PAPER ROCK
* PAPER PAPER
- PAPER SCISSORS

- SCISSORS ROCK
- SCISSORS PAPER
* SCISSORS SCISSORS

X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

1st answer: 9651
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFilename = "input"

	ROCK int = iota
	PAPER
	SCISSORS

	LOST = 0
	DRAW = 3
	WON  = 6
)

var sign2score = map[int]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

// maps a sign to what gets defeated
var winTable = map[int]int{
	ROCK:     SCISSORS,
	PAPER:    ROCK,
	SCISSORS: ROCK,
}

// maps a sign to what lose to
var loseTable = map[int]int{
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	SCISSORS: ROCK,
}

func match(opponent, player int) (score int) {
	switch {
	case opponent == player:
		score += DRAW
	case opponent == ROCK && player == SCISSORS:
		// defeat
		score += LOST
	case opponent == SCISSORS && player == ROCK:
		// win
		score += WON

	case opponent == ROCK && player == PAPER:
		// win
		score += WON
	case opponent == PAPER && player == ROCK:
		// defeat
		score += LOST

	case opponent == PAPER && player == SCISSORS:
		// win
		score += WON
	case opponent == SCISSORS && player == PAPER:
		// defeat
		score += LOST
	}

	score += sign2score[player]

	return score
}

func opponentHand(s string) int {
	switch s {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	default:
		panic("wrong hand value")
	}
}

func playerHand(s string) int {
	switch s {
	case "X":
		return ROCK
	case "Y":
		return PAPER
	case "Z":
		return SCISSORS
	default:
		panic("wrong player hand value")
	}
}

func run() error {
	fh, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer fh.Close()

	var score int

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		if len(values) != 2 {
			return fmt.Errorf("invalid line: %q (split gave more than 2 values", line)
		}

		opp := opponentHand(values[0])

		// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
		switch values[1] {
		case "X":
			hand := winTable[opp]
			score += (sign2score[hand] + LOST)
		case "Y":
			score += (sign2score[opp] + DRAW)
		case "Z":
			hand := loseTable[opp]
			score += (sign2score[hand] + WON)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// not 11866 nor 10497
	fmt.Printf("score: %d\n", score)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
