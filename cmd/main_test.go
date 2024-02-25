package main

import "testing"

type MockGitVersionCommand struct {
}

func (*MockGitVersionCommand) Output() ([]byte, error) {
	return []byte("git version 1.23.456\n"), nil
}

func TestGitVersion(t *testing.T) {
	origShellCommandFunc := shellCommandFunc
	defer func() { shellCommandFunc = origShellCommandFunc }()

	shellCommandFunc = func(name string, args ...string) CommandExecutor {
		wantName := "git"
		if name != wantName {
			t.Errorf("command name: got %q, want %q", name, wantName)
		}

		wantArg1 := "--version"
		if args[0] != wantArg1 {
			t.Errorf("command arg1: got %q, want %q", args[0], wantArg1)
		}

		return &MockGitVersionCommand{}
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
