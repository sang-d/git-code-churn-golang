package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	before := os.Args[1]
	after := os.Args[2]
	dir := os.Args[3]
	author := os.Args[4]

	log.Printf("Getting churn and new code for %s since %s until %s", author, after, before)

	commits := getCommits(before, after, dir, author)
	for _, commit := range commits {
		log.Println(commit)
	}
}

func execShellCommand(path string, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = path
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func getCommits(before, after, dir, author string) []string {
	command := "git"
	args := []string{"log", "--author", author, "--format=%h", "--no-abbrev",
		"--before", before, "--after", after, "--no-merges", "--reverse"}
	output, err := execShellCommand(dir, command, args...)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(output), "\n")
}
