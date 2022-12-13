package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

type Op int

const (
	Add Op = iota
	Mul
)

type Exp struct {
	operation Op
	// assume that if lhs or rhs is -1, it's "old"
	lhs int64
	rhs int64
}

type Monkey struct {
	id                int
	items_worry_level []int64

	// Operation shows how your worry level changes as that monkey inspects an item. (An operation
	// like new = old * 5 means that your worry level after the monkey inspected the item is five
	// times whatever your worry level was before inspection.)
	operationExpr Exp

	// Test shows how the monkey uses your worry level to decide where to throw an item next.
	test_divisible_by int64

	if_true_throw_to  int
	if_false_throw_to int

	items_inspected int64
}

func parseExp(input string) (Exp, error) {
	// assume string is of form "old * old" or "old + 3"
	re := regexp.MustCompile("^(.+) ([+*]) (.+)$")
	matches := re.FindStringSubmatch(input)
	if len(matches) != 4 {
		return Exp{}, fmt.Errorf("oh no, input formula nonsensical: %s", input)
	}
	e := Exp{}
	switch matches[2] {
	case "+":
		e.operation = Add
	case "*":
		e.operation = Mul
	default:
		return Exp{}, fmt.Errorf("oh no, operator nonsensical: %s", matches[2])
	}
	// assume that if lhs or rhs is -1, it's "old"
	if matches[1] == "old" {
		e.lhs = -1
	} else {
		lhs_i, err := strconv.Atoi(matches[1])
		if err != nil {
			return Exp{}, err
		}
		e.lhs = int64(lhs_i)
	}
	if matches[3] == "old" {
		e.rhs = -1
	} else {
		rhs_i, err := strconv.Atoi(matches[3])
		if err != nil {
			return Exp{}, err
		}
		e.rhs = int64(rhs_i)
	}
	return e, nil
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
			items_list := []int64{}
			for _, item_id_s := range items_split {
				item_id, err := strconv.Atoi(strings.TrimSpace(item_id_s))
				if err != nil {
					return []Monkey{}, err
				}
				items_list = append(items_list, int64(item_id))
			}
			m.items_worry_level = items_list
			state = Operation
		case Operation:
			operation_r := regexp.MustCompile("Operation: new = (.*)$")
			exp, err := parseExp(operation_r.FindStringSubmatch(i)[1])
			if err != nil {
				return []Monkey{}, err
			}
			m.operationExpr = exp
			state = Test
		case Test:
			test_r := regexp.MustCompile("Test: divisible by ([0-9]+)")
			test_str := test_r.FindStringSubmatch(i)[1]
			test_divisor, err := strconv.Atoi(strings.TrimSpace(test_str))
			if err != nil {
				return []Monkey{}, err
			}
			m.test_divisible_by = int64(test_divisor)
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

func (e Exp) String() string {
	op_s := "*"
	if e.operation == Add {
		op_s = "+"
	}
	lhs := "old"
	rhs := "old"
	if e.lhs > 0 {
		lhs = fmt.Sprint(e.lhs)
	}
	if e.rhs > 0 {
		rhs = fmt.Sprint(e.rhs)
	}
	return fmt.Sprintf("%s %s %s", lhs, op_s, rhs)
}

func (e Exp) eval(old int64) (answer int64, err error) {
	lhs := e.lhs
	rhs := e.rhs
	if lhs == -1 {
		lhs = old
	}
	if rhs == -1 {
		rhs = old
	}
	switch e.operation {
	case Mul:
		answer = lhs * rhs
	case Add:
		answer = lhs + rhs
	}

	if answer < 0 {
		// oh shit, overflow
		return 0, fmt.Errorf("Overflow! Answer = %d\n", answer)
	}
	return answer, nil
}

const ROUNDS = 10_000

func monkeyTurn(m int) (err error) {
	/// Monkey 0:
	///   Monkey inspects an item with a worry level of 79.
	///     Worry level is multiplied by 19 to 1501.
	///     Monkey gets bored with item. Worry level is divided by 3 to 500.
	///     Current worry level is not divisible by 23.
	///     Item with worry level 500 is thrown to monkey 3.
	fmt.Printf("Monkey %d:\n", m)

	me := &monkeys[m]
	my_items := me.items_worry_level

	// empty out our inventory - after this turn we won't have items
	me.items_worry_level = []int64{}

	for _, itm := range my_items {
		old_worry := itm
		fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", old_worry)
		me.items_inspected++
		new_worry, err := me.operationExpr.eval(old_worry)
		if err != nil {
			return err
		}
		fmt.Printf("    Worry level is \"%s\" to %d.\n", me.operationExpr, new_worry)
		if new_worry%me.test_divisible_by == 0 {
			fmt.Printf("    Current worry level is divisible by %d.\n", me.test_divisible_by)
			fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", new_worry, me.if_true_throw_to)
			monkeys[me.if_true_throw_to].items_worry_level = append(monkeys[me.if_true_throw_to].items_worry_level, new_worry)
		} else {
			fmt.Printf("    Current worry level is not divisible by %d.\n", me.test_divisible_by)
			fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", new_worry, me.if_false_throw_to)
			monkeys[me.if_false_throw_to].items_worry_level = append(monkeys[me.if_false_throw_to].items_worry_level, new_worry)
		}
	}
	return nil
}

func performAllRounds() (err error) {
	for round := 1; round <= ROUNDS; round++ {
		for m := range monkeys {
			err := monkeyTurn(m)
			if err != nil {
				return err
			}
		}

		fmt.Printf("\nAfter round %d, the monkeys are holding items with these worry levels:\n", round)
		for m := range monkeys {
			fmt.Printf("Monkey %d: %v\n", m, monkeys[m].items_worry_level)
		}
	}
	return nil
}

var monkeys []Monkey

func hottestMonkeys(monkeys []Monkey) []int64 {
	hot := []int64{}
	for m := range monkeys {
		hot = append(hot, monkeys[m].items_inspected)
	}
	// sort by descending order of hotness - note flipped ><><
	sort.Slice(hot, func(i, j int) bool { return hot[i] > hot[j] })
	return hot
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
	monkeys, err = parseMonkeyList(input)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", monkeys)

	err = performAllRounds()
	if err != nil {
		return err
	}
	for m := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", monkeys[m].id, monkeys[m].items_inspected)
	}

	hot := hottestMonkeys(monkeys)
	fmt.Printf("Hottest monkeys: %v\n", hot)

	monkey_business := hot[0] * hot[1]
	fmt.Printf("Monkey business: %d\n", monkey_business)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
