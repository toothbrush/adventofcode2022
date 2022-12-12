package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ParseState int

const (
	MonkeyIntro ParseState = iota
	Items
	Operation
	Test
	IfTrue
	IfFalse
)

type Monkey struct {
	id    int
	items []int

	operation string

	test_divisible_by int

	if_true_throw_to  int
	if_false_throw_to int
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
			intro_r := regexp.MustCompile("Monkey ([0-9]+):")
			id_string := intro_r.FindStringSubmatch(i)[1]
			id, err := strconv.Atoi(id_string)
			if err != nil {
				return []Monkey{}, err
			}
			m.id = id
			state = Items
		case Items:
			items_r := regexp.MustCompile("Starting items: ([0-9, ]+)")
			items_string := items_r.FindStringSubmatch(i)[1]
			items_split := strings.Split(items_string, ",")
			items_list := []int{}
			for _, item_id_s := range items_split {
				item_id, err := strconv.Atoi(strings.TrimSpace(item_id_s))
				if err != nil {
					return []Monkey{}, err
				}
				items_list = append(items_list, item_id)
			}
			m.items = items_list
			state = Operation
		case Operation:
			operation_r := regexp.MustCompile("Operation: new = (.*)$")
			m.operation = operation_r.FindStringSubmatch(i)[1]
			state = Test
		case Test:
			test_r := regexp.MustCompile("Test: divisible by ([0-9]+)")
			test_str := test_r.FindStringSubmatch(i)[1]
			test_divisor, err := strconv.Atoi(strings.TrimSpace(test_str))
			if err != nil {
				return []Monkey{}, err
			}
			m.test_divisible_by = test_divisor
			state = IfTrue
		case IfTrue:
			true_r := regexp.MustCompile("If true: throw to monkey ([0-9]+)")
			true_str := true_r.FindStringSubmatch(i)[1]
			m_id, err := strconv.Atoi(strings.TrimSpace(true_str))
			if err != nil {
				return []Monkey{}, err
			}
			m.if_true_throw_to = m_id
			state = IfFalse
		case IfFalse:
			false_r := regexp.MustCompile("If false: throw to monkey ([0-9]+)")
			false_str := false_r.FindStringSubmatch(i)[1]
			m_id, err := strconv.Atoi(strings.TrimSpace(false_str))
			if err != nil {
				return []Monkey{}, err
			}
			m.if_false_throw_to = m_id
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
