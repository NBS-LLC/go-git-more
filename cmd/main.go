package main

import (
	"fmt"
	"os/exec"
)

func GitVersion() (string, error) {
	cmd := exec.Command("git", "--version")
	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func main() {
	version, _ := GitVersion()
	fmt.Println(version)
}
