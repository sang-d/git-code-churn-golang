package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecShellCommandOutput(t *testing.T) {
	command := "ls"
	args := []string{"go.mod"}
	dir := "./"
	output, err := execShellCommand(dir, command, args...)
	assert.Nil(t, err)
	assert.Equal(t, string(output), "go.mod\n")
}

func TestGetCommits(t *testing.T) {
	// use this repo to test
	before, after, dir, author := "2020-07-23", "2020-07-20", "./", "sang.dinh@wizeline.com"
	commits := getCommits(before, after, dir, author)
	assert.Equal(t, "1d2e6e56cfb8c85eff1909dfc4eeefe8f3d99f33", commits[0])
}
