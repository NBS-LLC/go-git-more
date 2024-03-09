package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/mod/semver"
)

type commandExecutor interface {
	Output() ([]byte, error)
}

var shellCommandFunc = func(name string, arg ...string) commandExecutor {
	return exec.Command(name, arg...)
}

// GitVersion retrieves the version of git.
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

// GetTags retrieves a list of git tags sorted by version in descending order.
func GetTags() ([]string, error) {
	cmd := shellCommandFunc("git", "tag", "--list", "--sort=-version:refname")
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

// GetLatestTag returns the latest tag from the given list of tags.
func GetLatestTag(tags []string) string {
	semver.Sort(tags)
	return tags[len(tags) - 1]
}

func main() {
	version, _ := GitVersion()
	fmt.Println("Git Version:", version)

	tags, _ := GetTags()
	fmt.Println("Tags:", tags)
	fmt.Println("Latest Tag:", GetLatestTag(tags))
}
