package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type commandExecutor interface {
	Output() ([]byte, error)
}

var shellCommandFunc = func(name string, arg ...string) commandExecutor {
	return exec.Command(name, arg...)
}

func GitVersion() (string, error) {
	cmd := shellCommandFunc("git", "--version")
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out))[12:], nil
}

func GetTags() ([]string, error) {
	cmd := shellCommandFunc("git", "tag", "--list", "--sort=-version:refname")
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

func main() {
	version, _ := GitVersion()
	fmt.Println("Git Version:", version)

	tags, _ := GetTags()
	fmt.Println("Tags:", tags)
}
