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

type Inode struct {
	name     string
	size     uint32 // let's identify directories as size == 0 and children > 0.
	children []Inode
}

type FSState struct {
	cwd    []string // our current working directory,
	inodes []Inode  // filesystem contents.
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

func NewFSState() FSState {
	fs := FSState{}
	fs.cwd = []string{}
	fs.inodes = []Inode{}
	return fs
}

func (fs FSState) pwd() {
	fmt.Printf("/%s\n", strings.Join(fs.cwd, "/"))
}

func (fs *FSState) ls(output []string) error {
	// figure out what to do with FSState based on ls output
	return nil
}

func (fs *FSState) cd(dir string) error {
	// modify FSState as if the user changed directory

	fmt.Printf("Change directory to `%s`\n", dir)
	// special cases first:
	if dir == "/" {
		fs.cwd = []string{}
	} else if dir == ".." {
		// drop deepest directory
		if len(fs.cwd) == 0 {
			return fmt.Errorf("can't change directory up relative to /!")
		}
		fs.cwd = fs.cwd[0 : len(fs.cwd)-1]
	} else {
		fs.cwd = append(fs.cwd, dir)
	}
	fs.pwd()
	return nil
}

func (fs *FSState) executeCommand(cmd Command) (err error) {
	fmt.Printf("Executing `%s %s`\n", cmd.cmd, cmd.args)
	switch cmd.cmd {
	case "cd":
		err = fs.cd(cmd.args[0])
	case "ls":
		err = fs.ls(cmd.output)
	default:
		err = fmt.Errorf("unknown executable `%s`!", cmd.cmd)
	}
	return err
}

func (fs *FSState) executeCommands(cmds []Command) error {
	for _, cmd := range cmds {
		err := fs.executeCommand(cmd)
		if err != nil {
			return err
		}
	}
	return nil
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
			command, err := NewCommand(t)
			if err != nil {
				return err
			}
			commands = append(commands, command)
		} else {
			// it's the output of a command
			// by definition we're working on the last command added to the list
			if len(strings.TrimSpace(t)) > 0 { // skip blank lines
				last := &commands[len(commands)-1]
				last.output = append(last.output, t)
			}
		}
	}

	fs := NewFSState()
	fs.executeCommands(commands)

	fmt.Printf("%v\n", commands)
	return nil
}
