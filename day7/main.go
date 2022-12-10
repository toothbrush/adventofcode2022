package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type Command struct {
	cmd    string
	args   []string
	output []string
}

func NewCommand(cmdLine string) (Command, error) {
	split := strings.Split(cmdLine, " ")
	if split[0] != "$" {
		return Command{}, fmt.Errorf("oh scheisse command doesn't start with $!")
	}
	command := Command{}
	command.cmd = split[1]
	for _, arg := range split[2:] {
		command.args = append(command.args, arg)
	}
	command.output = make([]string, 0)
	return command, nil
}

func run() (err error) {
	s := bufio.NewScanner(os.Stdin)
	var t string

	var commands []Command
	commands = make([]Command, 0)

	for s.Scan() {
		t = s.Text()
		if strings.HasPrefix(t, "$") {
			// oh this is a command being executed
			// let's make a new slot for it
			fmt.Printf("executing: `%s`\n", t)
			command, err := NewCommand(t)
			if err != nil {
				return err
			}
			commands = append(commands, command)
		} else {
			// it's the output of a command
			// by definition we're working on the last command added to the list
			if len(strings.TrimSpace(t)) > 0 { // skip blank lines
				fmt.Printf("output: [%s]\n", t)
				last := &commands[len(commands)-1]
				last.output = append(last.output, t)
			}
		}

	}

	fmt.Printf("%v\n", commands)
	return nil
}
