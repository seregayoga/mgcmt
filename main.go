package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}

	commitMessage, err := parseBranchName(string(branch))
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("git", "commit", "-m", commitMessage).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func parseBranchName(branch string) (string, error) {
	branch = strings.TrimSpace(branch)
	r, _ := regexp.Compile("^([A-Za-z]+)-([0-9]+)-(.*)$")
	m := r.FindStringSubmatch(branch)
	if len(m) != 4 {
		return "", errors.New("Incorrect branch name, must starts with XXX-000-etc")
	}

	return fmt.Sprintf("%s-%s %s", m[1], m[2], strings.Replace(m[3], "-", " ", -1)), nil
}
