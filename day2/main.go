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

func main() {
	s := bufio.NewScanner(os.Stdin)

	var t string
	var line []string
	var challenge Move
	var response Move

	for s.Scan() {
		t = s.Text()
		line = strings.Split(t, " ")

		challenge = fromChallenge(line[0])
		response = fromResponse(line[1])

		fmt.Printf("challenge: %s\n", challenge)
		fmt.Printf("response:  %s\n", response)

	}
}
