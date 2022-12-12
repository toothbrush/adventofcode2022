package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type ParseState int

const (
	MonkeyIntro ParseState = iota
	Items
	Test
	IfTrue
	IfFalse
	MonkeyComplete
)

type Monkey struct {
	id int
}

func parseMonkeyList(input []string) ([]Monkey, error) {
	state := MonkeyIntro
	var m Monkey
	ms := []Monkey{}
	for _, i := range input {
		fmt.Printf("[%d] parsing: `%s`\n", state, i)
		switch state {
		case MonkeyIntro:
			m = Monkey{}
			intro := regexp.MustCompile("Monkey ([0-9]+):")
			id_string := intro.FindStringSubmatch(i)[1]
			id, err := strconv.Atoi(id_string)
			if err != nil {
				return []Monkey{}, err
			}
			m.id = id
			state = MonkeyComplete
		case Items:
		case Test:
		case IfTrue:
		case IfFalse:
		case MonkeyComplete:
			ms = append(ms, m)
			state = MonkeyIntro
		}
	}

	return ms, nil
}

func run() (err error) {
	fmt.Printf("welcome to monkeys\n")

	s := bufio.NewScanner(os.Stdin)
	var t string

	input := make([]string, 0)
	for s.Scan() {
		t = s.Text()
		if len(t) > 0 {
			fmt.Printf("%s\n", t)
			input = append(input, t)
		}
	}
	monkeys, err := parseMonkeyList(input)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", monkeys)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
