package dev08

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("MyShell> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		if strings.Contains(input, "|") {
			runPipeline(input, os.Stdin, os.Stdout)
		} else {
			err := runSingleCommand(input, os.Stdin, os.Stdout)
			if err != nil {
				return
			}
		}
	}
}

func runSingleCommand(input string, stdin io.Reader, stdout io.Writer) error {
	args := strings.Fields(input)
	if len(args) == 0 {
		return errors.New("len(args) == 0")
	}

	var cmd *exec.Cmd

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("Usage: cd <directory>")
		} else {
			err := os.Chdir(args[1])
			if err != nil {
				return err
			}
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		} else {
			fmt.Fprintln(stdout, dir)
		}
	case "echo":
		fmt.Fprintln(stdout, strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			return errors.New("Usage: kill <process_id>")
		} else {
			processID := args[1]
			cmd = exec.Command("kill", processID)
		}
	default:
		cmd = exec.Command(args[0], args[1:]...)
		cmd.Stdin = stdin
		cmd.Stdout = stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			return err
		}
		err = cmd.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}

func runPipeline(input string, stdin io.Reader, stdout io.Writer) {
	commands := strings.Split(input, "|")

	for _, command := range commands {
		err := runSingleCommand(strings.TrimSpace(command), stdin, stdout)
		if err != nil {
			fmt.Println("error with single command", err)
		}
	}
}
