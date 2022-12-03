package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Move int8

const (
	Rock Move = iota
	Paper
	Scissors
)

type PairOutcome int8

const (
	Lose PairOutcome = iota
	Draw
	Win
)

func (m Move) String() string {
	switch m {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	}
	panic("unreasonable move provided")
}

func fromChallenge(c string) Move {
	switch c {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}
	panic(fmt.Sprintf("unreasonable challenge provided: %s", c))
}

func fromResponse(c string) Move {
	switch c {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	}
	panic(fmt.Sprintf("unreasonable response provided: %s", c))
}

func (c Move) losesTo(r Move) PairOutcome {
	switch c {
	case Rock:
		switch r {
		case Rock:
			return Draw
		case Paper:
			return Win
		case Scissors:
			return Lose
		}
	case Paper:
		switch r {
		case Rock:
			return Lose
		case Paper:
			return Draw
		case Scissors:
			return Win
		}
	case Scissors:
		switch r {
		case Rock:
			return Win
		case Paper:
			return Lose
		case Scissors:
			return Draw
		}
	}
	panic("impossible")
}

func pointsForRound(challenge Move, response Move) (mypoints int64) {
	mypoints = 0

	// did i win?
	switch challenge.losesTo(response) {
	case Win:
		mypoints += 6
	case Draw:
		mypoints += 3
	case Lose:
		mypoints += 0
	}

	// add points for the response i used
	switch response {
	case Rock:
		mypoints += 1
	case Paper:
		mypoints += 2
	case Scissors:
		mypoints += 3
	}

	return mypoints
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	var t string
	var line []string
	var challenge Move
	var response Move

	var thisRoundPoints int64
	var totalPoints int64 = 0

	for s.Scan() {
		t = s.Text()
		line = strings.Split(t, " ")

		challenge = fromChallenge(line[0])
		response = fromResponse(line[1])

		fmt.Printf("challenge: %s\n", challenge)
		fmt.Printf("response:  %s\n", response)

		thisRoundPoints = pointsForRound(challenge, response)
		fmt.Printf("points for round = %d\n", thisRoundPoints)
		totalPoints += thisRoundPoints
	}

	fmt.Printf("total points: %d\n", totalPoints)
}
