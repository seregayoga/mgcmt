package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}
	separatedBranch := strings.Split(string(branch), "-")

	taskName := strings.Join(separatedBranch[:2], "-")
	comment := strings.Join(separatedBranch[2:], " ")
	comment = strings.TrimSpace(comment)

	commitMessage := fmt.Sprintf("%s %s", taskName, comment)

	err = exec.Command("git", "commit", "-m", commitMessage).Run()
	if err != nil {
		log.Fatal(err)
	}
}
