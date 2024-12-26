package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	cmd := exec.Command("cmd", "/C", strings.Join(args, " "))

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command
	return cmd.Run()
}
