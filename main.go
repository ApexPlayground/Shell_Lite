package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type GOOS int

const (
	Linux GOOS = iota
	Windows
	Darwin
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("-> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Execution of the input
		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove the newLine character.
	input = strings.TrimSpace(input)

	// Split the input into command and arguments
	args := strings.Split(input, " ")

	// Handle the 'cd' command separately
	if args[0] == "cd" {
		if len(args) < 2 {
			return fmt.Errorf("cd: missing argument")
		}
		return os.Chdir(args[1])
	}

	// For other commands, execute them directly
	cmd, err := execCommand(args)
	if err != nil || cmd == nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Execute the command
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func execCommand(args []string) (*exec.Cmd, error) {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("cmd", "/C", strings.Join(args, " ")), nil
	case "darwin", "linux":
		return exec.Command("/bin/sh", "-c", strings.Join(args, " ")), nil
	default:
		return nil, fmt.Errorf("OS not supported")
	}
}
