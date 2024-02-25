package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type commandExecutor interface {
	Output() ([]byte, error)
}

type shellCommandFunc func(name string, arg ...string) commandExecutor

func shellCommand(name string, arg ...string) commandExecutor {
	return exec.Command(name, arg...)
}

func GitVersion(s shellCommandFunc) (string, error) {
	cmd := s("git", "--version")
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out))[12:], nil
}

func main() {
	version, _ := GitVersion(shellCommand)
	fmt.Println("Git Version:", version)
}
