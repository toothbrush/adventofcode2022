package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	children map[string]Inode
}

type FSState struct {
	cwd  []string // our current working directory,
	root Inode    // filesystem contents. named inodes.
}

func (i Inode) sizeOrDir() string {
	if i.size == 0 {
		return "dir"
	} else {
		return fmt.Sprintf("file=%d", i.size)
	}
}

func PrintInode(depth int, inode Inode) string {
	s := ""
	indent := strings.Repeat("  ", depth)
	s += fmt.Sprintf("%s- %s (%s)\n", indent, inode.name, inode.sizeOrDir())
	for _, v := range inode.children {
		s += PrintInode(depth+1, v)
	}
	return s
}

func (fs FSState) String() string {
	return PrintInode(0, fs.root)
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
	fs.root = NewDirectory("/")
	return fs
}

func NewDirectory(name string) Inode {
	kids := make(map[string]Inode)
	return Inode{name: name, size: 0, children: kids}
}

func NewFile(name string, size uint32) Inode {
	kids := make(map[string]Inode)
	return Inode{name: name, size: size, children: kids}
}

func pwd(cwd []string) string {
	return fmt.Sprintf("/%s", strings.Join(cwd, "/"))
}

func addChildTo(inode Inode, cwd []string, child Inode) (Inode, error) {
	if len(cwd) == 0 {
		// we have recursed sufficiently, add it here
		c := inode.children
		if c == nil {
			fmt.Printf("r u nil?? %v\n", inode)
			c = make(map[string]Inode)
			inode.children = c
		} else {
			fmt.Printf("exists??  %v\n", inode)
		}
		c[child.name] = child
		fmt.Printf("after adding = %v\n", inode)
		return inode, nil
	} else {
		c := inode.children
		if c == nil {
			fmt.Printf("r u nil?? %v\n", inode)
			c = make(map[string]Inode)
			inode.children = c
		} else {
			fmt.Printf("exists??  %v\n", inode)
		}
		// very suspect are we detaching things??
		inode := c[cwd[0]]
		new, err := addChildTo(inode, cwd[1:], child)
		if err != nil {
			return Inode{}, err
		}
		c[cwd[0]] = new
		return inode, nil
	}
}

func (fs *FSState) addInode(inode Inode) error {
	fmt.Printf("Adding '%v' to %s\n", inode, pwd(fs.cwd))
	newRoot, err := addChildTo(fs.root, fs.cwd, inode)
	if err != nil {
		return err
	}
	fs.root = newRoot
	return nil
}

func (fs *FSState) ls(output []string) (err error) {
	// figure out what to do with FSState based on ls output
	for _, lsLine := range output {
		split := strings.Split(lsLine, " ")
		var new_inode Inode
		if split[0] == "dir" {
			new_inode = NewDirectory(split[1])
		} else {
			atoi, err := strconv.Atoi(split[0])
			if err != nil {
				return err
			}
			sz := uint32(atoi)
			new_inode = NewFile(split[1], sz)
		}
		err = fs.addInode(new_inode)
	}
	return err
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
	fmt.Printf("cwd: %s\n", pwd(fs.cwd))
	fmt.Printf("%s\n", fs)
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
