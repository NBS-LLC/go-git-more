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

	shellCommandCalled := false
	shellCommandFunc = func(name string, args ...string) commandExecutor {
		shellCommandCalled = true
		assert.Equal("git", name, "command name")
		assert.Len(args, 1, "command args")
		assert.Equal("--version", args[0], "1st command arg")
		return &MockCommandExecutor{output: "git version 1.23.456\n"}
	}

	version, err := GitVersion()
	if assert.NoError(err) {
		assert.Equal("1.23.456", version, "version string")
	}

	assert.True(shellCommandCalled, "shell command called")
}

func TestGetTags(t *testing.T) {
	assert := assert.New(t)

	origShellCommandFunc := shellCommandFunc
	defer func() { shellCommandFunc = origShellCommandFunc }()

	shellCommandCalled := false
	shellCommandFunc = func(name string, args ...string) commandExecutor {
		shellCommandCalled = true
		assert.Equal("git", name, "command name")
		assert.Contains(args, "--sort=-version:refname", "command args")
		return &MockCommandExecutor{output: "v1.0.0\nv0.1.0\nv0.0.1\n"}
	}

	expectedTags := []string{"v1.0.0", "v0.1.0", "v0.0.1"}
	actualTags, err := GetTags()
	if assert.NoError(err) {
		assert.Equal(expectedTags, actualTags, "tags")
	}

	assert.True(shellCommandCalled, "shell command called")
}

func TestGetLatestTag(t *testing.T) {
	tags := []string{"v0.1.0", "v1.0.0", "v1.1.0", "v0.2.1"}
	assert.Equal(t, "v1.1.0", GetLatestTag(tags))
}
