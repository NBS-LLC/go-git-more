package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type CommandExecutor interface {
	Output() ([]byte, error)
}

func shellCommand(name string, arg ...string) CommandExecutor {
	return exec.Command(name, arg...)
}

var shellCommandFunc = shellCommand

func GitVersion() (string, error) {
	cmd := shellCommandFunc("git", "--version")
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func main() {
	version, _ := GitVersion()
	fmt.Println(version)
}
