package main

import "testing"

type MockCommandExecutor struct {
	output string
}

func (m *MockCommandExecutor) Output() ([]byte, error) {
	return []byte(m.output), nil
}

func TestGitVersion(t *testing.T) {
	origShellCommandFunc := shellCommandFunc
	defer func() { shellCommandFunc = origShellCommandFunc }()

	shellCommandFunc = func(name string, args ...string) commandExecutor {
		wantName := "git"
		if name != wantName {
			t.Errorf("command name: got %q, want %q", name, wantName)
		}

		wantArg1 := "--version"
		if args[0] != wantArg1 {
			t.Errorf("command arg1: got %q, want %q", args[0], wantArg1)
		}

		return &MockCommandExecutor{output: "git version 1.23.456\n"}
	}

	got, err := GitVersion()
	want := "1.23.456"

	if err != nil {
		t.Fatalf("got an error")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
