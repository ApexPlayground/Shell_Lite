package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GOOS int

const (
	Linux GOOS = iota
	Windows
	Darwin
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Shell Lite")

	// Input field and output box
	inputField := widget.NewEntry()
	inputField.SetPlaceHolder("Enter a command...")

	outputLabel := widget.NewLabel("Output will appear here.")

	// Execute button
	executeButton := widget.NewButton("Execute", func() {
		command := inputField.Text
		output, err := execInput(command)
		if err != nil {
			outputLabel.SetText(fmt.Sprintf("Error: %s", err))
		} else {
			outputLabel.SetText(output)
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Command Executor"),
		inputField,
		executeButton,
		outputLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
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
