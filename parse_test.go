package main

import "testing"

func TestParseBranchNameOK(t *testing.T) {
	branch := "PRJ-123-some-feature"

	commitMessage, err := parseBranchName(branch)
	if err != nil {
		t.Error(err)
	}

	if commitMessage != "PRJ-123 some feature" {
		t.Errorf("Incorrect commit message '%s' from branch '%s'", commitMessage, branch)
	}
}

func TestParseBranchNameInvalid(t *testing.T) {
	branches := [...]string{
		"",
		"master",
		"wrong-branch-name",
	}

	for _, branch := range branches {
		commitMessage, err := parseBranchName(branch)
		if err == nil {
			t.Error("parseBranchName() should return error")
		}

		if commitMessage != "" {
			t.Error("parseBranchName() should return empty string")
		}
	}
}
