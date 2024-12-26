package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("Shell Lite - Command Executor")
	fmt.Println("Type 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)

		// Exit condition
		if input == "exit" {
			fmt.Println("Exiting Shell Lite. Goodbye!")
			break
		}

		// Execute the command
		output, err := execInput(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Println(output)
		}
	}
}

// execInput processes the command input and executes it
func execInput(input string) (string, error) {
	// Remove newline and split into arguments
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")

	// Handle 'cd' command separately
	if args[0] == "cd" {
		if len(args) < 2 {
			return "", fmt.Errorf("cd: missing argument")
		}
		err := os.Chdir(args[1])
		if err != nil {
			return "", err
		}
		return "Changed directory successfully.", nil
	}

	// Execute other commands
	cmd, err := execCommand(args)
	if err != nil || cmd == nil {
		return "", err
	}

	outputBytes, err := cmd.CombinedOutput()
	return string(outputBytes), err
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
