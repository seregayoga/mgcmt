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

	commitMessage, err := parseBranchName(branch)
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("git", "commit", "-m", commitMessage).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func parseBranchName(branch []byte) (task string, err error) {
	r, _ := regexp.Compile("^([A-Za-z]+)-([0-9]+)-(.*)$")
	m := r.FindStringSubmatch(string(branch))
	if len(m) != 4 {
		return "", errors.New("Incorrect branch name, must starts with XXX-000-etc")
	}

	return fmt.Sprintf("%s-%s %s", m[1], m[2], strings.Replace(m[3], "-", " ", -1)), nil
}
