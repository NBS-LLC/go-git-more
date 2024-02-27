package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCommandExecutor struct {
	output string
}

func (m *MockCommandExecutor) Output() ([]byte, error) {
	return []byte(m.output), nil
}

func TestGitVersion(t *testing.T) {
	assert := assert.New(t)

	origShellCommandFunc := shellCommandFunc
	defer func() { shellCommandFunc = origShellCommandFunc }()

	shellCommandFunc = func(name string, args ...string) commandExecutor {
		assert.Equal("git", name, "command name")
		assert.Len(args, 1, "command args")
		assert.Equal("--version", args[0], "1st command arg")
		return &MockCommandExecutor{output: "git version 1.23.456\n"}
	}

	version, err := GitVersion()
	if assert.NoError(err) {
		assert.Equal("1.23.456", version, "version string")
	}
}
