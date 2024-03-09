package main

import (
	"fmt"
	"os/exec"
	"regexp"
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

	re := regexp.MustCompile(`\b\d+\.\d+\.\d+\b`)
	version := re.FindString(string(out))
	return version, nil
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
	fmt.Println("Most Recent Tag:", tags[0])
}
